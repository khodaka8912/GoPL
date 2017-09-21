package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type testStruct struct {
	Map     map[string]bool
	Float   float32
	Boolean bool
	Slice   []string
}

func main() {
	a := testStruct{
		map[string]bool{
			"Key1": true,
			"Key2": false,
		}, 123.456, true, []string{"str1", "str2", "str3"},
	}
	bytes, err := Marshal(a)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Println("encoded = " + string(bytes))
	b := testStruct{}
	if err = json.Unmarshal(bytes, &b); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("unmarshaled = %v\n", b)
}

// Marshal encodes a Go value in json form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an json representation of v.
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

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

	case reflect.Array, reflect.Slice: // [value ...]
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // {"name":value, ...}
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(buf, "%q:", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // {key:value,...}
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	default: // complex, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
