// Copy from https://github.com/thinkeridea/go-extend

package elstr

import (
	"strings"
	"unicode/utf8"
)

// Replace strings but return []byte
func ReplaceToBytes(s, old, new string, n int) []byte {
	if old == new || n == 0 {
		return []byte(s) // avoid allocation
	}
	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return []byte(s) // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}
	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return t[0:w]
}

// Replace strings but return []byte
// Dangerous!!
// This function is very dangerous (but really fast)
// If the string was written to your source code directly,
// Please don't use this function. Because it will cause a series error.
// If you want to use this function, Please check bytes.UnsafeToBytes
func UnsafeReplaceToBytes(s, old, new string, n int) []byte {
	if old == new || n == 0 {
		return UnsafeToBytes(s) // avoid allocation
	}
	// Compute number of replacements.
	if m := strings.Count(s, old); m == 0 {
		return UnsafeToBytes(s) // avoid allocation
	} else if n < 0 || m < n {
		n = m
	}
	// Apply replacements to buffer.
	t := make([]byte, len(s)+n*(len(new)-len(old)))
	w := 0
	start := 0
	for i := 0; i < n; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(s[start:])
				j += wid
			}
		} else {
			j += strings.Index(s[start:], old)
		}
		w += copy(t[w:], s[start:j])
		w += copy(t[w:], new)
		start = j + len(old)
	}
	w += copy(t[w:], s[start:])
	return t[0:w]
}

// Repeat the string but return []byte
// It panics if count is negative or if
// the result of (len(s) * count) overflows.
func RepeatToBytes(s string, count int) []byte {
	// Since we cannot return an error on overflow,
	// we should panic if the repeat will generate
	// an overflow.
	// See Issue golang.org/issue/16237
	if count < 0 {
		panic("strings: negative Repeat count")
	} else if count > 0 && len(s)*count/count != len(s) {
		panic("strings: Repeat count causes overflow")
	}
	b := make([]byte, len(s)*count)
	bp := copy(b, s)
	for bp < len(b) {
		copy(b[bp:], b[:bp])
		bp *= 2
	}
	return b
}

// Join strings but return []byte
func JoinToBytes(a []string, sep string) []byte {
	switch len(a) {
	case 0:
		return []byte{}
	case 1:
		return []byte(a[0])
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}
	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return b
}
