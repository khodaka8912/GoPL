package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Circle struct {
	name string
	tail *Circle
}

func main() {
	var a, b, c, d Circle
	a.tail = &b
	b.tail = &c
	c.tail = &a
	d.tail = &Circle{}

	fmt.Println(HasCirculation(a))
	fmt.Println(HasCirculation(d))
}

func HasCirculation(v interface{}) bool {
	seen := make(map[value]bool)
	return hasCirculation(reflect.ValueOf(v), seen)
}

func hasCirculation(v reflect.Value, seen map[value]bool) bool {
	if !v.IsValid() {
		return false
	}
	if v.CanAddr() {
		ptr := unsafe.Pointer(v.UnsafeAddr())
		val := value{ptr, v.Type()}
		if seen[val] {
			return true
		}
		seen[val] = true

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
