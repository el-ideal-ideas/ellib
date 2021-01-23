// This module contains some helper functions for array and slice.
// For example:
// var s = []string{"も", "ず", "く"}
// StrReverse(s)
// fmt.Println(s)
// >> []string{"く", "ず", "も"}

package arr

// Reverse slice
func Reverse(s []interface{}) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func StrReverse(s []string) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func UintReverse(s []uint) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Uint8Reverse(s []uint8) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Uint16Reverse(s []uint16) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Uint32Reverse(s []uint32) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Uint64Reverse(s []uint64) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func IntReverse(s []int) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Int8Reverse(s []int8) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Int16Reverse(s []int16) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Int32Reverse(s []int32) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Int64Reverse(s []int64) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func ByteReverse(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func RuneReverse(s []rune) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Float32Reverse(s []float32) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}

// Reverse slice
func Float64Reverse(s []float64) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		li := length - i - 1
		s[i], s[li] = s[li], s[i]
	}
}
