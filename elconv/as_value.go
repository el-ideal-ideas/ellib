package elconv

import (
	"reflect"
)

// If `v` is a pointer, return the value.
func AsValue(v interface{}) interface{} {
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Ptr {
		return r.Elem().Interface()
	} else {
		return r
	}
}

// If `r` is a pointer, return the value.
func AsValueRef(r reflect.Value) reflect.Value {
	if r.Kind() == reflect.Ptr {
		return r.Elem()
	} else {
		return r
	}
}
