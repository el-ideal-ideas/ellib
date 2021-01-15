// UnsafeToBytes is the fastest way to convert string to []byte, but this function is unsafe.
// []byte() < ToBytes() <<<<< UnsafeToBytes()
// BenchmarkDefaultToBytes-8        7494451               156 ns/op             768 B/op          1 allocs/op
// BenchmarkStrToBytes-8            8400860               143 ns/op             768 B/op          1 allocs/op
// BenchmarkStrUnsafeToBytes-8     1000000000               0.320 ns/op           0 B/op          0 allocs/op

package str

import (
	"unsafe"
)


// Convert string to []byte (faster than ToBytes)
// Dangerous!!
// This function is very dangerous (but really fast)
// If the string was written to your source code directly,
// Please don't use this function. Because it will cause a series error.
// ** string literals can't be changed from program **
// s := `{"k":"v"}`
// b := UnsafeToBytes(s)
// b[3] = 'k' // unexpected fault address 0x1118180
// If you use string literals and try change the result, The error will happen.
func UnsafeToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// Convert string to []byte (faster than []byte("some string"))
func ToBytes(s string) []byte {
	buf := make([]byte, len(s))
	copy(buf, s)
	return buf
}

// Convert []byte to string (faster than string([]byte{1, 2, 3}))
// Important:
// If you changed `s`, the result will be changed too.
// string data in golang is immutable, But this function use unsafe pointer.
// When you change the slice, the string will be changed too.
func UnsafeToString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}
