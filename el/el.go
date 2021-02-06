package el

import (
	"github.com/el-ideal-ideas/ellib/elconv"
	"github.com/el-ideal-ideas/ellib/eldev"
	"github.com/el-ideal-ideas/ellib/elref"
)

var IsNil = elref.IsNil
var IsEmpty = elref.IsEmpty
var Type = elref.Type
var IsIn = elref.IsIn

var IsStruct = elref.IsStruct
var IsInterface = elref.IsInterface
var IsMap = elref.IsMap
var IsArray = elref.IsArray
var IsSlice = elref.IsSlice
var IsArrayOrSlice = elref.IsArrayOrSlice
var IsBool = elref.IsBool
var IsInt = elref.IsInt
var IsInt8 = elref.IsInt8
var IsInt16 = elref.IsInt16
var IsInt32 = elref.IsInt32
var IsInt64 = elref.IsInt64
var IsUint = elref.IsUint
var IsUint8 = elref.IsUint8
var IsUint16 = elref.IsUint16
var IsUint32 = elref.IsUint32
var IsUint64 = elref.IsUint64
var IsUintptr = elref.IsUintptr
var IsFloat32 = elref.IsFloat32
var IsFloat64 = elref.IsFloat64
var IsComplex64 = elref.IsComplex64
var IsComplex128 = elref.IsComplex128
var IsChan = elref.IsChan
var IsFunc = elref.IsFunc
var IsPtr = elref.IsPtr
var IsString = elref.IsString
var IsUnsafePointer = elref.IsUnsafePointer
var IsNumber = elref.IsNumber

var Bool = elconv.AsBool
var Float32 = elconv.AsFloat32
var Float64 = elconv.AsFloat64
var Int = elconv.AsInt
var Int64 = elconv.AsInt64
var String = elconv.AsStr
var Uint = elconv.AsUint
var Uint64 = elconv.AsUint64
var Value = elconv.AsValue
var Pointer = elconv.AsPtr
var Slice = elconv.AsSlice

var StructFields = elref.GetStructFields
var StructValues = elref.GetStructValues

var PanicIfErr = eldev.PanicIfErr

// IF evaluates a condition, if true returns the first parameter otherwise the second
func IF(condition bool, first interface{}, second interface{}) interface{} {
	if condition {
		return first
	} else {
		return second
	}
}

// NotNil check `v` parameter, if `v` is not nil, return it, otherwise return the default value.
func NotNil(v interface{}, def interface{}) interface{} {
	if elref.IsNil(v) {
		return def
	} else {
		return v
	}
}

// If any items in `v` is true, return true. Otherwise, return false.
// python has same build-in function
func Any(v ...bool) bool {
	for i := range v {
		if v[i] {
			return true
		}
	}
	return false
}

// If any items in `v` is false, return false. Otherwise, return true.
// python has same build-in function
func All(v ...bool) bool {
	for i := range v {
		if !v[i] {
			return false
		}
	}
	return true
}

// FirstNonNil returns the first non nil parameter.
// If all parameter is nil, return nil.
func FirstNonNil(values ...interface{}) interface{} {
	for _, value := range values {
		if !elref.IsNil(value) {
			return value
		}
	}
	return nil
}
