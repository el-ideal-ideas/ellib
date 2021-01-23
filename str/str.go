// This module contains some helper functions for string data.

package str

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// -------- these constants are from python's string module --------

// Whitespace -- a string containing all ASCII whitespace
// AsciiLowercase -- a string containing all ASCII lowercase letters
// AsciiUppercase -- a string containing all ASCII uppercase letters
// AsciiLetters -- a string containing all ASCII letters
// Digits -- a string containing all ASCII decimal digits
// HexDigits -- a string containing all ASCII hexadecimal digits
// OctDigits -- a string containing all ASCII octal digits
// PUNCTUATION -- a string containing all ASCII punctuation characters
// Printable -- a string containing all ASCII characters considered printable
const Whitespace = " \t\n\r\v\f"
const AsciiLowercase = "abcdefghijklmnopqrstuvwxyz"
const AsciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const AsciiLetters = AsciiLowercase + AsciiUppercase
const Digits = "0123456789"
const HexDigits = Digits + "abcdef" + "ABCDEF"
const OctDigits = "01234567"
const PUNCTUATION = `!"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`"
const Printable = Digits + AsciiLetters + PUNCTUATION + Whitespace

// -------- utility methods like python's str class --------

// Return True if the string is an alphabetic string, False otherwise.
// A string is alpha-numeric if all characters in the string are alpha-numeric and
// there is at least one character in the string.
func IsAlNum(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

// Return True if the string is an alphabetic string, False otherwise.
// A string is alphabetic if all characters in the string are alphabetic and there
// is at least one character in the string.
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// Return True if the string is a digit string, False otherwise.
// A string is a digit string if all characters in the string are digits and there
// is at least one character in the string.
func IsDigit(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// Return True if all characters in the string are ASCII, False otherwise.
// ASCII characters have code points in the range U+0000-U+007F.
// Empty string is ASCII too.
func IsASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// Return True if the string is an uppercase string, False otherwise.
// A string is uppercase if all cased characters in the string are uppercase and
// there is at least one cased character in the string.
func IsUpper(s string) bool {
	if s == "" {
		return false
	}
	flag := false
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			return false
		} else if r >= 'A' && r <= 'Z' {
			flag = true
		}
	}
	return flag
}

// Return True if the string is a lowercase string, False otherwise.
// A string is lowercase if all cased characters in the string are lowercase and
// there is at least one cased character in the string.
func IsLower(s string) bool {
	if s == "" {
		return false
	}
	flag := false
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return false
		} else if r >= 'a' && r <= 'z' {
			flag = true
		}
	}
	return flag
}

// -------- other utilities for strings --------

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

// Replace strings (faster than strings.Replace)
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
