package elref

import (
	"reflect"
)

// If `r` is a pointer, return the value.
func reflectPtr(r reflect.Value) reflect.Value {
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r
}

// If `v` is a struct, return true.
func IsStruct(v interface{}) bool {
	r := reflectPtr(reflect.ValueOf(v))
	return r.Kind() == reflect.Struct
}

// If `v` is a interface, return true.
func IsInterface(val interface{}) bool {
	r := reflectPtr(reflect.ValueOf(val))
	return r.Kind() == reflect.Invalid
}

// If `v` is nil return true.
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface:
		if rv.IsNil() {
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
	data := reflect.ValueOf(v)
	switch data.Kind() {
	case reflect.String, reflect.Array:
		return data.Len() == 0
	case reflect.Map, reflect.Slice:
		return data.Len() == 0 || data.IsNil()
	case reflect.Bool:
		return !data.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return data.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return data.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return data.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return data.IsNil()
	}
	return reflect.DeepEqual(v, reflect.Zero(data.Type()).Interface())
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

// If `v` is array, return true. Otherwise return false.
func IsArray(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Array
}

// If `v` is slice, return true. Otherwise return false.
func IsSlice(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Slice
}

// If `v` is array or slice, return true. Otherwise return false.
func IsArrayOrSlice(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Slice || reflect.ValueOf(v).Kind() == reflect.Array
}

// If `v` is map, return true. Otherwise return false.
func IsMap(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Map
}

// Returns a string of variable type
func Type(v interface{}) string {
	return reflect.TypeOf(v).String()
}
