package elref

import (
	"reflect"
)

// If `v` is nil return true.
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface:
		if r.IsNil() {
			return true
		}
	}
	return false
}

// If `v` is empty value, return true.
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.String, reflect.Array:
		return r.Len() == 0
	case reflect.Map, reflect.Slice:
		return r.IsNil() || r.Len() == 0
	case reflect.Bool:
		return !r.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return r.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return r.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return r.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return r.IsNil()
	}
	return reflect.DeepEqual(v, reflect.Zero(r.Type()).Interface())
}

// If `needle` is in `haystack` return true.
// supports array, slice and map.
func IsIn(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("[IsIn]haystack type must be array, slice or map")
	}
	return false
}

// Returns a string of variable type
func Type(v interface{}) string {
	return reflect.TypeOf(v).String()
}
