package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

type test struct {
	Map       map[string]bool
	Float     float32
	Boolean   bool
	Complex   complex128
	Interface fmt.Stringer
}

func main() {
	enc := NewEncoder(os.Stdout)
	a := test{
		map[string]bool{
			"Key1": true,
			"Key2": false,
		}, 123.456, true, complex(9.9, -0.5), nil,
	}
	if err := enc.Encode(a); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	zero := test{}
	if err := enc.Encode(zero); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(v interface{}) error {
	if err := encode(e.w, reflect.ValueOf(v)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(e.w, "\n"); err != nil {
		return err
	}
	return nil
}

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(w io.Writer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprint(w, "nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(w, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(w, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(w, "%q", v.String())

	case reflect.Ptr:
		return encode(w, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		fmt.Fprint(w, "(")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprint(w, " ")
			}
			if err := encode(w, v.Index(i)); err != nil {
				return err
			}
		}
		fmt.Fprint(w, ")")

	case reflect.Struct: // ((name value) ...)
		fmt.Fprint(w, "(")
		for i := 0; i < v.NumField(); i++ {
			if isZeroValue(v.Field(i)) {
				continue
			}
			if i > 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprintf(w, "(%s ", v.Type().Field(i).Name)
			if err := encode(w, v.Field(i)); err != nil {
				return err
			}
			fmt.Fprint(w, ")")
		}
		fmt.Fprint(w, ")")

	case reflect.Map: // ((key value) ...)
		fmt.Fprint(w, "(")
		for i, key := range v.MapKeys() {
			if i > 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, "(")
			if err := encode(w, key); err != nil {
				return err
			}
			fmt.Fprint(w, " ")
			if err := encode(w, v.MapIndex(key)); err != nil {
				return err
			}
			fmt.Fprint(w, ")")
		}
		fmt.Fprint(w, ")")
	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(w, "t")
		} else {
			fmt.Fprint(w, "nil")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(w, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(w, "#C(%f %f)", real(c), imag(c))
	case reflect.Interface:
		fmt.Fprint(w, "(")
		t := v.Type()
		fmt.Fprintf(w, `"%s.%s" `, t.PkgPath(), t.Name())
		if err := encode(w, v.Elem()); err != nil {
			return err
		}
		fmt.Fprint(w, ")")
	default: // chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Complex128, reflect.Complex64:
		c := v.Complex()
		return real(c) == 0 && imag(c) == 0
	}
	return false
}
