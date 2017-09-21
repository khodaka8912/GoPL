package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Request struct {
	Name    string `http:"n"`
	Age     int
	Email   string   `http:"mail", type:"email"`
	Zip     string   `type:"zip"`
	Credits []string `http:"cr", type:"credit"`
}

func main() {
	valid := "http://test.valid?n=testName&age=10&mail=mail@test.com&zip=1234567&cr=123456789012356&cr=6543210987654321"
	u, _ := url.Parse(valid)
	req := Request{}
	if err := Unpack(&http.Request{URL: u}, &req); err != nil {
		fmt.Errorf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println(req)
	invalid := "http://test.invalid?n=invalid&mail=invalid&zip=0&cr=none"
	u, _ = url.Parse(invalid)
	req2 := Request{}
	err := Unpack(&http.Request{URL: u}, &req2)
	if err == nil {
		fmt.Errorf("cannot detect conditoin error")
		os.Exit(1)
	}
	fmt.Println(err)

}

type field struct {
	v         reflect.Value
	condition *regexp.Regexp
}

var email = regexp.MustCompile(`^[a-zA-Z0-9\-.]+@[a-zA-Z0-9\-.]+$`)
var credit = regexp.MustCompile(`^[0-9]{16}$`)
var zip = regexp.MustCompile(`^[0-9]{7}$`)

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]field)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		paramType := tag.Get("type")
		var condition *regexp.Regexp
		switch paramType {
		case "email":
			condition = email
		case "credit":
			condition = credit
		case "zip":
			condition = zip
		}
		fields[name] = field{v.Field(i), condition}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.v.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.v.Kind() == reflect.Slice {
				elem := reflect.New(f.v.Type().Elem()).Elem()
				if err := populate(field{elem, f.condition}, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.v.Set(reflect.Append(f.v, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(f field, value string) error {
	if f.condition != nil && !f.condition.MatchString(value) {
		return fmt.Errorf("%q does not satisfy the condition", value)
	}
	switch f.v.Kind() {
	case reflect.String:
		f.v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		f.v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f.v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", f.v.Type())
	}
	return nil
}
