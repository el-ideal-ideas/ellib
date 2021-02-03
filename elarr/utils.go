package elarr

import (
	"github.com/el-ideal-ideas/ellib/elconv"
	"github.com/el-ideal-ideas/ellib/elstr"
	"reflect"
)

// Sum of numbers in `v`
func Sum(v []interface{}) float64 {
	var sum float64
	for i := 0; i < len(v); i++ {
		sum += elconv.AsFloat64(v[i])
	}
	return sum
}

// Sum of numbers in `v`
func SumInt(v []int) int {
	var sum int
	for i := 0; i < len(v); i++ {
		sum += v[i]
	}
	return sum
}

// Average of numbers in `v`
func Average(v []interface{}) float64 {
	return Sum(v) / float64(len(v))
}

// Join can join items in array&slice as string.
func Join(v []interface{}, sep string) string {
	return elstr.Join(ToStrings(v), sep)
}

// If `v` is a empty slice or array, return true. Otherwise, return false.
func IsEmpty(v interface{}) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		return val.Len() == 0
	} else {
		return false
	}
}
