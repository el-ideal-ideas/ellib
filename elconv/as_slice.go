package elconv

import (
	"reflect"
)

// convert `v` to a slice
func AsSlice(v interface{}) []interface{} {
	if v == nil {
		return []interface{}{}
	}
	r := AsValueRef(reflect.ValueOf(v))
	switch r.Kind() {
	case reflect.Slice, reflect.Array:
		res := make([]interface{}, r.Len())
		for i := 0; i < r.Len(); i++ {
			res[i] = r.Index(i).Interface()
		}
		return res
	case reflect.Map:
		res := make([]interface{}, r.Len())
		for i, k := range r.MapKeys() {
			res[i] = k.Interface()
		}
		return res
	case reflect.Struct:
		t := r.Type()
		res := make([]string, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			res[i] = t.Field(i).Name
		}
	}
	return []interface{}{v}
}
