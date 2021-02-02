// This module contains some helper functions for array and slice.

package elarr

import (
	"github.com/el-ideal-ideas/ellib/elconv"
	"github.com/el-ideal-ideas/ellib/elref"
	"reflect"
)

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
// Maybe slow.
var IsSame = IsSameInter

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
// Maybe slow.
func IsSameInter(data ...interface{}) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check type
	for i := 0; i < count; i++ {
		if !elref.IsArrayOrSlice(data[i]) {
			return false
		}
	}
	// check length
	length := reflect.ValueOf(data[0]).Len()
	for i := 1; i < count; i++ {
		if reflect.ValueOf(data[i]).Len() != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		v := reflect.ValueOf(data[0])
		first := v.Index(i).Interface()
		for j := 1; j < count; j++ {
			switch first.(type) {
			case int, int8, int16, int32, int64:
				if elconv.AsInt64(first) != elconv.AsInt64(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case uint, uint8, uint16, uint32, uint64:
				if elconv.AsUint64(first) != elconv.AsUint64(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case float32:
				if elconv.AsFloat32(first) != elconv.AsFloat32(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case float64:
				if elconv.AsFloat64(first) != elconv.AsFloat64(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case string:
				if elconv.AsStr(first) != elconv.AsStr(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			default:
				if !reflect.DeepEqual(first, reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameInterSlice(data ...[]interface{}) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			switch first.(type) {
			case int, int8, int16, int32, int64:
				if elconv.AsInt64(first) != elconv.AsInt64(data[j][i]) {
					return false
				}
			case uint, uint8, uint16, uint32, uint64:
				if elconv.AsUint64(first) != elconv.AsUint64(data[j][i]) {
					return false
				}
			case float32:
				if elconv.AsFloat32(first) != elconv.AsFloat32(reflect.ValueOf(data[j]).Index(i).Interface()) {
					return false
				}
			case float64:
				if elconv.AsFloat64(first) != elconv.AsFloat64(data[j][i]) {
					return false
				}
			case string:
				if elconv.AsStr(first) != elconv.AsStr(data[j][i]) {
					return false
				}
			default:
				if !reflect.DeepEqual(first, data[j][i]) {
					return false
				}
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameStr(data ...[]string) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameInt(data ...[]int) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameInt8(data ...[]int8) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameInt16(data ...[]int16) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameInt32(data ...[]int32) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameInt64(data ...[]int64) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameUint(data ...[]uint) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameUint8(data ...[]uint8) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameUint16(data ...[]uint16) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameUint32(data ...[]uint32) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameUint64(data ...[]uint64) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameRune(data ...[]rune) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameByte(data ...[]byte) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameFloat32(data ...[]float32) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameFloat64(data ...[]float64) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}
