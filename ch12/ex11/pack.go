package main

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strings"
)

type Request struct {
	Name     string
	page     int `http:"p"`
	children []string
	student  bool
}

func main() {
	u, _ := url.Parse("http://test.url")
	req := Request{"testName", 5, []string{"child1", "child2"}, true}
	if err := Pack(u, req); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Println(u)
}

func Pack(url *url.URL, ptr interface{}) error {
	q := url.Query()
	v := reflect.ValueOf(ptr)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		f := v.Field(i)
		if f.Kind() == reflect.Slice {
			len := f.Len()
			for i := 0; i < len; i++ {
				elm := f.Index(i)
				value, err := toString(elm)
				if err != nil {
					return err
				}
				q.Add(name, value)
			}
		} else {
			value, err := toString(f)
			if err != nil {
				return err
			}
			q.Set(name, value)
		}
	}
	url.RawQuery = q.Encode()
	return nil
}

func toString(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil

	case reflect.Int:
		return fmt.Sprintf("%d", v.Int()), nil

	case reflect.Bool:
		if v.Bool() {
			return "true", nil
		} else {
			return "false", nil
		}

	default:
		return "", fmt.Errorf("unsupported kind %s", v.Type())
	}
}

//!-populate
