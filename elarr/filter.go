package elarr

import "reflect"

// Filter will return a new slice containing only the elements
// that return true from the condition.
// supports array, slice.
// this function use reflect, will more slow than other Filter functions.
func Filter(data interface{}, condition func(interface{}) bool) []interface{} {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		res := make([]interface{}, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			d := v.Index(i).Interface()
			if condition(d) {
				res = append(res, d)
			}
		}
		return res
	default:
		return []interface{}{}
	}
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterInter(data []interface{}, condition func(interface{}) bool) []interface{} {
	res := make([]interface{}, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterStr(data []string, condition func(string) bool) []string {
	res := make([]string, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterInt(data []int, condition func(int) bool) []int {
	res := make([]int, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterInt8(data []int8, condition func(int8) bool) []int8 {
	res := make([]int8, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterInt16(data []int16, condition func(int16) bool) []int16 {
	res := make([]int16, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterInt32(data []int32, condition func(int32) bool) []int32 {
	res := make([]int32, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterInt64(data []int64, condition func(int64) bool) []int64 {
	res := make([]int64, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterUint(data []uint, condition func(uint) bool) []uint {
	res := make([]uint, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterUint8(data []uint8, condition func(uint8) bool) []uint8 {
	res := make([]uint8, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterUint16(data []uint16, condition func(uint16) bool) []uint16 {
	res := make([]uint16, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterUint32(data []uint32, condition func(uint32) bool) []uint32 {
	res := make([]uint32, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterUint64(data []uint64, condition func(uint64) bool) []uint64 {
	res := make([]uint64, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterRune(data []rune, condition func(rune) bool) []rune {
	res := make([]rune, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterByte(data []byte, condition func(byte) bool) []byte {
	res := make([]byte, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterFloat32(data []float32, condition func(float32) bool) []float32 {
	res := make([]float32, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

// Filter will return a new slice containing only the elements
// that return true from the condition.
func FilterFloat64(data []float64, condition func(float64) bool) []float64 {
	res := make([]float64, 0, len(data))
	for i := 0; i < len(data); i++ {
		if condition(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}
