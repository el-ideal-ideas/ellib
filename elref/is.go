package elref

import (
	"github.com/el-ideal-ideas/ellib/elconv"
	"reflect"
)

// If `v` is a struct, return true. Otherwise return false.
func IsStruct(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Struct
}

// If `v` is a interface, return true. Otherwise return false.
func IsInterface(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Invalid
}

// If `v` is map, return true. Otherwise return false.
func IsMap(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Map
}

// If `v` is array, return true. Otherwise return false.
func IsArray(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Array
}

// If `v` is slice, return true. Otherwise return false.
func IsSlice(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Slice
}

// If `v` is array or slice, return true. Otherwise return false.
func IsArrayOrSlice(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Slice || reflect.ValueOf(v).Kind() == reflect.Array
}

// If `v` is a bool, return true. Otherwise return false.
func IsBool(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Bool
}

// If `v` is a Int, return true. Otherwise return false.
func IsInt(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Int
}

// If `v` is a Int8, return true. Otherwise return false.
func IsInt8(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Int8
}

// If `v` is a Int16, return true. Otherwise return false.
func IsInt16(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Int16
}

// If `v` is a Int32, return true. Otherwise return false.
func IsInt32(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Int32
}

// If `v` is a Int64, return true. Otherwise return false.
func IsInt64(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Int64
}

// If `v` is a Uint, return true. Otherwise return false.
func IsUint(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Uint
}

// If `v` is a Uint8, return true. Otherwise return false.
func IsUint8(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Uint8
}

// If `v` is a Uint16, return true. Otherwise return false.
func IsUint16(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Uint16
}

// If `v` is a Uint32, return true. Otherwise return false.
func IsUint32(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Uint32
}

// If `v` is a Uint64, return true. Otherwise return false.
func IsUint64(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Uint64
}

// If `v` is a Uintptr, return true. Otherwise return false.
func IsUintptr(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Uintptr
}

// If `v` is a Float32, return true. Otherwise return false.
func IsFloat32(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Float32
}

// If `v` is a Float64, return true. Otherwise return false.
func IsFloat64(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Float64
}

// If `v` is a Complex64, return true. Otherwise return false.
func IsComplex64(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Complex64
}

// If `v` is a Complex128, return true. Otherwise return false.
func IsComplex128(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Complex128
}

// If `v` is a Chan, return true. Otherwise return false.
func IsChan(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Chan
}

// If `v` is a Func, return true. Otherwise return false.
func IsFunc(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.Func
}

// If `v` is a Ptr, return true. Otherwise return false.
func IsPtr(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Ptr
}

// If `v` is a String, return true. Otherwise return false.
func IsString(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.String
}

// If `v` is a String, return true. Otherwise return false.
func IsUnsafePointer(v interface{}) bool {
	r := elconv.AsValueRef(reflect.ValueOf(v))
	return r.Kind() == reflect.UnsafePointer
}
