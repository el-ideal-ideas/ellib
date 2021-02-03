// This module contains some helper functions for array and slice.
// For example:
// var s = []string{"も", "ず", "く"}
// ReverseStr(s)
// fmt.Println(s)
// >> []string{"く", "ず", "も"}

package elarr

// Reverse slice
func Reverse(s []interface{}) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseStr(s []string) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseUint(s []uint) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseUint8(s []uint8) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseUint16(s []uint16) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseUint32(s []uint32) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseUint64(s []uint64) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseInt(s []int) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseInt8(s []int8) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseInt16(s []int16) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseInt32(s []int32) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseInt64(s []int64) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseByte(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseRune(s []rune) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseFloat32(s []float32) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseFloat64(s []float64) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ReverseBool(s []bool) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}
