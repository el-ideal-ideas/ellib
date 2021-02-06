// This module contains some helper functions for string data.

package elstr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var IsEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)

// alias of fmt.S
var Format = fmt.Sprintf

// Reverse string
func Reverse(s string) string {
	var start, size, end int
	buf := make([]byte, len(s))
	for end < len(s) {
		_, size = utf8.DecodeRuneInString(s[start:])
		end = start + size
		copy(buf[len(buf)-end:], s[start:end])
		start = end
	}
	return UnsafeToString(buf)
}

// Reverse string (only supports ascii strings, faster than Reverse function)
func ReverseASCII(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return UnsafeToString(b)
}

// Replace string (faster than strings.Replace)
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string {
	return UnsafeToString(UnsafeReplaceToBytes(s, old, new, n))
}

// Repeat strings (faster than strings.Repeat)
// It panics if count is negative or if
// the result of (len(s) * count) overflows.
func Repeat(s string, count int) string {
	return UnsafeToString(RepeatToBytes(s, count))
}

// Join strings (faster than strings.Join)
func Join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	}
	return UnsafeToString(JoinToBytes(a, sep))
}

// JoinInts join []int as string.
func JoinInts(ints []int, sep string) (res string) {
	length := len(ints)
	if length == 0 {
		return
	}
	var sb strings.Builder
	for _, num := range ints {
		sb.WriteString(strconv.Itoa(num))
		if length--; length > 0 {
			sb.WriteString(sep)
		}
	}
	res = sb.String()
	return
}

// Copy strings for GC.
// s := "A really big string"
// c := Copy(s[1:10])
// If you only want to use a part of `s`, You can copy the part. So GC can collect `s`
func Copy(src string) string {
	buf := make([]byte, len(src))
	copy(buf, src)
	return UnsafeToString(buf)
}

// Get a part of string (supports rune).
func SubString(s string, start, length int) string {
	return RuneSubString(s, start, length)
}

// Split string by multiple separators. will clear empty string node.
func SplitMulti(s string, sep ...string) (ss []string) {
	if len(sep) == 0 {
		ss = append(ss, s)
		return
	}
	if s = strings.TrimSpace(s); s == "" {
		return
	}
	if len(sep) == 1 {
		for _, val := range strings.Split(s, sep[0]) {
			if val = strings.TrimSpace(val); val != "" {
				ss = append(ss, val)
			}
		}
	} else {
		for i := 1; i < len(sep); i++ {
			s = Replace(s, sep[i], sep[0], -1)
		}
		return SplitMulti(s, sep[0])
	}
	return
}

// Replaces replace multi strings
//
// 	pairs: {old1: new1, old2: new2, ...}
//
// Can also use:
// 	strings.NewReplacer("old1", "new1", "old2", "new2").Replace(str)
func Replaces(str string, pairs map[string]string) string {
	ss := make([]string, len(pairs)*2)
	for old, newVal := range pairs {
		ss = append(ss, old, newVal)
	}
	return strings.NewReplacer(ss...).Replace(str)
}

// Convert hex string to int.
// Hex2Dec("0xff")
func Hex2Dec(str string) (int64, error) {
	start := 0
	if len(str) > 2 && str[0:2] == "0x" {
		start = 2
	}
	return strconv.ParseInt(str[start:], 16, 0)
}

// If `s` is a valid format as email address, return true.
func IsEmailFormat(s string) bool {
	if len(s) == 0 {
		return false
	}
	return IsEmailRegex.MatchString(s)
}

// If all characters in `s` is a Chinese character, return true.
func IsAllChineseChar(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}

// If `s` is palindrome, return true.
func IsPalindrome(s string) bool {
	r := []rune(s)
	length := len(r)
	for i := 0; i < length/2; i++ {
		if r[i] != r[length-i-1] {
			return false
		}
	}
	return true
}

// get the longest common substring of `s1` and `s2`.
// https://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Longest_common_substring#Go
func LCS(s1 string, s2 string) string {
	var m = make([][]int, 1+len(s1))
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, 1+len(s2))
	}
	longest := 0
	xLongest := 0
	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				m[x][y] = m[x-1][y-1] + 1
				if m[x][y] > longest {
					longest = m[x][y]
					xLongest = x
				}
			}
		}
	}
	return s1[xLongest-longest : xLongest]
}
