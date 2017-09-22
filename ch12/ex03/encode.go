package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
)

func main() {
	a := struct {
		Map       map[string]bool
		Float     float32
		Boolean   bool
		Complex   complex128
		Interface fmt.Stringer
		Test      interface{}
	}{
		map[string]bool{
			"Key1": true,
			"Key2": false,
		}, 123.456, true, complex(9.9, -0.5), bytes.NewBufferString("BufStr"),
		[]string{"a"},
	}
	bytes, err := Marshal(a)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
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
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("t")
		} else {
			buf.WriteString("nil")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%f %f)", real(c), imag(c))
	case reflect.Interface:
		buf.WriteByte('(')
		t := v.Elem().Type()
		fmt.Fprintf(buf, "%q ", t.String())
		if err := encode(buf, v.Elem()); err != nil {
			return err
		}
		buf.WriteByte(')')
	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
