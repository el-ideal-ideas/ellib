package elconv

import "reflect"

// get the pointer of `v`
func AsPtr(v interface{}) uintptr {
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r.Pointer()
}

// get the pointer of `r`
func AsPtrRef(r reflect.Value) uintptr {
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r.Pointer()
}
