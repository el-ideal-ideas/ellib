package elstr

import (
	"strings"
	"testing"
)

// BenchmarkStringsReplace-8        4449276               268 ns/op             192 B/op          2 allocs/op
// BenchmarkStrReplace-8            5381660               224 ns/op              96 B/op          1 allocs/op

var s = Repeat("もっちもっちにゃんにゃん!", 20)
var arr = strings.Split(s, "")

func BenchmarkStringsReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Replace(s, "もっち", "もち", -1)
	}
}

func BenchmarkStrReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Replace(s, "もっち", "もち", -1)
	}
}

// BenchmarkStringsRepeat-8         5327054               222 ns/op             768 B/op          1 allocs/op
// BenchmarkStrRepeat-8             5843203               204 ns/op             768 B/op          1 allocs/op

func BenchmarkStringsRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Repeat("もっちもっちにゃんにゃん!", 20)
	}
}

func BenchmarkStrRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Repeat("もっちもっちにゃんにゃん!", 20)
	}
}

// BenchmarkStringsJoin-8            378478              3174 ns/op            1024 B/op          1 allocs/op
// BenchmarkStrJoin-8                421634              2821 ns/op            1024 B/op          1 allocs/op

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(arr, ",")
	}
}

func BenchmarkStrJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Join(arr, ",")
	}
}

// BenchmarkDefaultToBytes-8        7494451               156 ns/op             768 B/op          1 allocs/op
// BenchmarkStrToBytes-8            8400860               143 ns/op             768 B/op          1 allocs/op
// BenchmarkStrUnsafeToBytes-8     1000000000               0.320 ns/op           0 B/op          0 allocs/op

func BenchmarkDefaultToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkStrToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToBytes(s)
	}
}

func BenchmarkStrUnsafeToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UnsafeToBytes(s)
	}
}
