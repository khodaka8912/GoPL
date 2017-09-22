package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Circle struct {
	name   string
	next   *Circle
	before *Circle
}

func main() {
	var a, b, c, d, e Circle
	a.next = &b
	b.next = &c
	c.next = &a
	d.next = &e
	d.before = &e

	fmt.Println(HasCirculation(a))
	fmt.Println(HasCirculation(d))
}

func HasCirculation(v interface{}) bool {
	seen := make([]value, 0)
	return hasCirculation(reflect.ValueOf(v), seen)
}

func hasCirculation(v reflect.Value, seen []value) bool {
	if !v.IsValid() {
		return false
	}
	if v.CanAddr() {
		ptr := unsafe.Pointer(v.UnsafeAddr())
		val := value{ptr, v.Type()}
		for _, s := range seen {
			if val == s {
				return true
			}
		}
		seen = append(seen, val)
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return hasCirculation(v.Elem(), seen)

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if hasCirculation(v.Field(i), seen) {
				return true
			}
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if hasCirculation(v.Index(i), seen) {
				return true
			}
		}

	case reflect.Map:
		for _, k := range v.MapKeys() {
			if hasCirculation(v.MapIndex(k), seen) {
				return true
			}
		}
	}
	return false
}

type value struct {
	ptr unsafe.Pointer
	t   reflect.Type
}
