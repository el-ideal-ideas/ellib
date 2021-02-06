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
// supports array, slice, struct and map.
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
	case reflect.Struct:
		t := val.Type()
		for i := 0; i < t.NumField(); i++ {
			if reflect.DeepEqual(needle, t.Field(i).Name) {
				return true
			}
		}
	default:
		panic("[IsIn]haystack type must be array, slice, struct or map")
	}
	return false
}

// Returns a string of variable type
func Type(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// Get fields in struct.
func GetStructFields(v interface{}) []string {
	if !IsStruct(v) {
		panic("[GetStructFields] v must be a struct")
	}
	t := reflect.ValueOf(v).Type()
	res := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		res[i] = t.Field(i).Name
	}
	return res
}

// Get values in struct.
func GetStructValues(v interface{}) []interface{} {
	if !IsStruct(v) {
		panic("[GetStructValues] v must be a struct")
	}
	r := reflect.ValueOf(v)
	t := r.Type()
	res := make([]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		res[i] = r.Field(i).Interface()
	}
	return res
}
