package elarr

import (
	"github.com/el-ideal-ideas/ellib/elconv"
	"github.com/el-ideal-ideas/ellib/elrand"
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

// Max of numbers in `v`
func Max(v []interface{}) float64 {
	var max = elconv.AsFloat64(v[0])
	for i := 1; i < len(v); i++ {
		if t := elconv.AsFloat64(v[i]); max < t {
			max = t
		}
	}
	return max
}

// Max of numbers in `v`
func MaxInt(v []int) int {
	var max = v[0]
	for i := 1; i < len(v); i++ {
		if max < v[i] {
			max = v[i]
		}
	}
	return max
}

// Min of numbers in `v`
func Min(v []interface{}) float64 {
	var min = elconv.AsFloat64(v[0])
	for i := 1; i < len(v); i++ {
		if t := elconv.AsFloat64(v[i]); min > t {
			min = t
		}
	}
	return min
}

// Min of numbers in `v`
func MinInt(v []int) int {
	var min = v[0]
	for i := 1; i < len(v); i++ {
		if min > v[i] {
			min = v[i]
		}
	}
	return min
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
// If `v` is not a slice or array, return true.
func IsEmpty(v interface{}) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		return val.Len() == 0
	} else {
		return true
	}
}

var Shuffle = elrand.Shuffle
