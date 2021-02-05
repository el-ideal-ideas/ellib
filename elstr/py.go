package elstr

import (
	"strings"
	"unicode"
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

// Return True if the string is an alpha-numeric string, False otherwise.
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

// Strip returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func Strip(s string) string {
	return strings.TrimSpace(s)
}
