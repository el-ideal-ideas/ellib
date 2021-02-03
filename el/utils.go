package el

import "github.com/el-ideal-ideas/ellib/elref"

// IF evaluates a condition, if true returns the first parameter otherwise the second
func IF(condition bool, first interface{}, second interface{}) interface{} {
	if condition {
		return first
	} else {
		return second
	}
}

// NotNil check `v` parameter, if `v` is not nil, return it, otherwise return the default value.
func NotNil(v interface{}, def interface{}) interface{} {
	if elref.IsNil(v) {
		return def
	} else {
		return v
	}
}

// If any items in `v` is true, return true. Otherwise, return false.
func Any(v ...bool) bool {
	for i := range v {
		if v[i] {
			return true
		}
	}
	return false
}

// If any items in `v` is false, return false. Otherwise, return true.
func All(v ...bool) bool {
	for i := range v {
		if !v[i] {
			return false
		}
	}
	return true
}

// FirstNonNil returns the first non nil parameter.
// If all parameter is nil, return nil.
func FirstNonNil(values ...interface{}) interface{} {
	for _, value := range values {
		if !elref.IsNil(value) {
			return value
		}
	}
	return nil
}
