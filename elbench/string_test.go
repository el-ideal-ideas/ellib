// 指针类型比字符串类型速度要快。
// *string is faster than string.
// 文字列本体よりもポインターを使用した方が高速である。
// BenchmarkString10-8              4102509               266 ns/op              90 B/op          0 allocs/op
// BenchmarkStringPtr10-8           5560388               219 ns/op              46 B/op          0 allocs/op
// BenchmarkString50-8              5112308               278 ns/op              96 B/op          0 allocs/op
// BenchmarkStringPtr50-8           5775960               226 ns/op              60 B/op          0 allocs/op
// BenchmarkString1000-8            5209543               244 ns/op              95 B/op          0 allocs/op
// BenchmarkStringPtr1000-8         5748008               222 ns/op              60 B/op          0 allocs/op

package elbench

import (
	"strings"
	"testing"
)

var storage = make(map[int]string)
var storagePtr = make(map[int]*string)
var data10 = strings.Repeat("もっちもっちにゃんにゃん！", 10)
var data50 = strings.Repeat("もっちもっちにゃんにゃん！", 50)
var data1000 = strings.Repeat("もっちもっちにゃんにゃん！", 1000)

func BenchmarkString10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set(i, data10)
	}
}

func BenchmarkStringPtr10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setPtr(i, &data10)
	}
}

func BenchmarkString50(b *testing.B) {
	storage = make(map[int]string)
	for i := 0; i < b.N; i++ {
		set(i, data50)
	}
}

func BenchmarkStringPtr50(b *testing.B) {
	storagePtr = make(map[int]*string)
	for i := 0; i < b.N; i++ {
		setPtr(i, &data50)
	}
}

func BenchmarkString1000(b *testing.B) {
	storage = make(map[int]string)
	for i := 0; i < b.N; i++ {
		set(i, data1000)
	}
}

func BenchmarkStringPtr1000(b *testing.B) {
	storagePtr = make(map[int]*string)
	for i := 0; i < b.N; i++ {
		setPtr(i, &data1000)
	}
}

func set(i int, s string) {
	storage[i] = s
}

func setPtr(i int, s *string) {
	storagePtr[i] = s
}
