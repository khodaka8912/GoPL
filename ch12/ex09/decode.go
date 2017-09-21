package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"text/scanner"
)

func main() {
	sexpr := `((Map (("Key1" "Value1") ("Key2" "Value2"))) (Integer 123) (Boolean t) (Array ("a" "b" "c"))`
	dec := NewDecoder(bytes.NewReader([]byte(sexpr)))
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		fmt.Print(tokenString(t) + " ")
	}
}

type Token interface{}

type Symbol struct {
	Value string
}
type String struct {
	Value string
}
type Int struct {
	Value int
}
type StartList struct{}

type EndList struct{}

type Decoder struct {
	lex *lexer
}

func tokenString(token Token) string {
	switch t := token.(type) {
	case Symbol:
		return t.Value
	case String:
		return fmt.Sprintf("\"%s\"", t.Value)
	case Int:
		return fmt.Sprintf("%d", t.Value)
	case StartList:
		return "("
	case EndList:
		return ")"
	}
	panic(fmt.Sprintf("unexpected token %q", token))
}

func Unmarshal(data []byte, out interface{}) error {
	return NewDecoder(bytes.NewReader(data)).Decode(out)
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	lex.next() // get the first token
	return &Decoder{lex}
}

func (dec *Decoder) Decode(out interface{}) (err error) {
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.lex.scan.Position, x)
		}
	}()
	read(dec.lex, reflect.ValueOf(out).Elem())
	return nil
}

func (dec *Decoder) Token() (Token, error) {
	switch dec.lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names
		s := dec.lex.text()
		if s == "nil" {
			// No token for nil value
		}
		dec.lex.next()
		return Symbol{s}, nil
	case scanner.String:
		s, err := strconv.Unquote(dec.lex.text())
		dec.lex.next()
		if err != nil {
			return nil, err
		}
		return String{s}, nil
	case scanner.Int:
		i, _ := strconv.Atoi(dec.lex.text()) // NOTE: ignoring errors
		dec.lex.next()
		return Int{i}, nil
	case scanner.Float:
		//f, _ := strconv.ParseFloat(dec.lex.text(), 64) // NOTE: ignoring errors
		//dec.lex.next()
		//return Float{f}
	case '(':
		dec.lex.next()
		return StartList{}, nil
	case ')':
		dec.lex.next()
		return EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	}
	panic(fmt.Sprintf("unexpected token %q", dec.lex.text()))
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

// The read function is a decoder for a small subset of well-formed
// S-expressions.  For brevity of our example, it takes many dubious
// shortcuts.
//
// The parser assumes
// - that the S-expression input is well-formed; it does no error checking.
// - that the S-expression input corresponds to the type of the variable.
// - that all numbers in the input are non-negative decimal integers.
// - that all keys in ((key value) ...) struct syntax are unquoted symbols.
// - that the input does not contain dotted lists such as (1 2 . 3).
// - that the input does not contain Lisp reader macros such 'x and #'x.
//
// The reflection logic assumes
// - that v is always a variable of the appropriate type for the
//   S-expression value.  For example, v must not be a boolean,
//   interface, channel, or function, and if v is an array, the input
//   must have the correct number of elements.
// - that v in the top-level call to read has the zero value of its
//   type and doesn't need clearing.
// - that if v is a numeric variable, it is a signed integer.
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil", "t", and struct field names
		s := lex.text()
		if s == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
		if s == "t" {
			v.SetBool(true)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64) // NOTE: ignoring errors
		v.SetFloat(f)
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

//!-read

//!+readlist
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
