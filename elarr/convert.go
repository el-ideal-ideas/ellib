package elarr

import (
	"math"
	"strconv"
	"unicode/utf8"
)

// convert slice data.
func ToStrings(arr []interface{}) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = v.(string)
	}
	return res
}

// convert slice data.
func ToInts(arr []interface{}) []int {
	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = v.(int)
	}
	return res
}

// convert slice data.
func ToInt8s(arr []interface{}) []int8 {
	res := make([]int8, len(arr))
	for i, v := range arr {
		res[i] = v.(int8)
	}
	return res
}

// convert slice data.
func ToInt16s(arr []interface{}) []int16 {
	res := make([]int16, len(arr))
	for i, v := range arr {
		res[i] = v.(int16)
	}
	return res
}

// convert slice data.
func ToInt32s(arr []interface{}) []int32 {
	res := make([]int32, len(arr))
	for i, v := range arr {
		res[i] = v.(int32)
	}
	return res
}

// convert slice data.
func ToInt64s(arr []interface{}) []int64 {
	res := make([]int64, len(arr))
	for i, v := range arr {
		res[i] = v.(int64)
	}
	return res
}

// convert slice data.
func ToUints(arr []interface{}) []uint {
	res := make([]uint, len(arr))
	for i, v := range arr {
		res[i] = v.(uint)
	}
	return res
}

// convert slice data.
func ToUint8s(arr []interface{}) []uint8 {
	res := make([]uint8, len(arr))
	for i, v := range arr {
		res[i] = v.(uint8)
	}
	return res
}

// convert slice data.
func ToUint16s(arr []interface{}) []uint16 {
	res := make([]uint16, len(arr))
	for i, v := range arr {
		res[i] = v.(uint16)
	}
	return res
}

// convert slice data.
func ToUint32s(arr []interface{}) []uint32 {
	res := make([]uint32, len(arr))
	for i, v := range arr {
		res[i] = v.(uint32)
	}
	return res
}

// convert slice data.
func ToUint64s(arr []interface{}) []uint64 {
	res := make([]uint64, len(arr))
	for i, v := range arr {
		res[i] = v.(uint64)
	}
	return res
}

// convert slice data.
func ToBytes(arr []interface{}) []byte {
	res := make([]byte, len(arr))
	for i, v := range arr {
		res[i] = v.(byte)
	}
	return res
}

// convert slice data.
func ToRunes(arr []interface{}) []rune {
	res := make([]rune, len(arr))
	for i, v := range arr {
		res[i] = v.(rune)
	}
	return res
}

// convert slice data.
func ToFloat32s(arr []interface{}) []float32 {
	res := make([]float32, len(arr))
	for i, v := range arr {
		res[i] = v.(float32)
	}
	return res
}

// convert slice data.
func ToFloat64s(arr []interface{}) []float64 {
	res := make([]float64, len(arr))
	for i, v := range arr {
		res[i] = v.(float64)
	}
	return res
}

// convert []int to []string.
func IntsToStrings(arr []int) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(v)
	}
	return res
}

// convert []int8 to []string.
func Int8sToStrings(arr []int8) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []int16 to []string.
func Int16sToStrings(arr []int16) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []int32 to []string.
func Int32sToStrings(arr []int32) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []int64 to []string.
func Int64sToStrings(arr []int64) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []uint to []string.
func UintsToStrings(arr []uint) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []uint8 to []string.
func Uint8sToStrings(arr []uint8) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []uint16 to []string.
func Uint16sToStrings(arr []uint16) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []uint32 to []string.
func Uint32sToStrings(arr []uint32) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []uint64 to []string.
func Uint64sToStrings(arr []uint64) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []rune to []string.
func RunesToStrings(arr []rune) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = string(v)
	}
	return res
}

// convert []byte to []string.
func BytesToStrings(arr []byte) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.Itoa(int(v))
	}
	return res
}

// convert []float32 to []string use strconv.FormatFloat.
//
// FormatFloat converts the floating-point number f to a string,
// according to the format fmt and precision prec. It rounds the
// result assuming that the original was obtained from a floating-point
// value of bitSize bits (32 for float32, 64 for float64).
//
// The format fmt is one of
// 'b' (-ddddp±ddd, a binary exponent),
// 'e' (-d.dddde±dd, a decimal exponent),
// 'E' (-d.ddddE±dd, a decimal exponent),
// 'f' (-ddd.dddd, no exponent),
// 'g' ('e' for large exponents, 'f' otherwise),
// 'G' ('E' for large exponents, 'f' otherwise),
// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
//
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', 'G', 'x', and 'X' formats.
// For 'e', 'E', 'f', 'x', and 'X', it is the number of digits after the decimal point.
// For 'g' and 'G' it is the maximum number of significant digits (trailing
// zeros are removed).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func Float32sToStrings(arr []float32, fmt byte, prec, bitSize int) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.FormatFloat(float64(v), fmt, prec, bitSize)
	}
	return res
}

// convert []float64 to []string use strconv.FormatFloat.
//
// FormatFloat converts the floating-point number f to a string,
// according to the format fmt and precision prec. It rounds the
// result assuming that the original was obtained from a floating-point
// value of bitSize bits (32 for float32, 64 for float64).
//
// The format fmt is one of
// 'b' (-ddddp±ddd, a binary exponent),
// 'e' (-d.dddde±dd, a decimal exponent),
// 'E' (-d.ddddE±dd, a decimal exponent),
// 'f' (-ddd.dddd, no exponent),
// 'g' ('e' for large exponents, 'f' otherwise),
// 'G' ('E' for large exponents, 'f' otherwise),
// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
//
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', 'G', 'x', and 'X' formats.
// For 'e', 'E', 'f', 'x', and 'X', it is the number of digits after the decimal point.
// For 'g' and 'G' it is the maximum number of significant digits (trailing
// zeros are removed).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func Float64sToStrings(arr []float64, fmt byte, prec, bitSize int) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = strconv.FormatFloat(v, fmt, prec, bitSize)
	}
	return res
}

// convert []int to []float64
func IntsToFloat64s(arr []int) []float64 {
	res := make([]float64, len(arr))
	for i, v := range arr {
		res[i] = float64(v)
	}
	return res
}

// convert []float64 to []int
func Float64sToInts(arr []float64) []int {
	res := make([]int, len(arr))
	for i, v := range arr {
		if v <= math.MaxInt64 && v >= math.MinInt64 {
			res[i] = int(v)
		} else if v > 0 {
			res[i] = math.MaxInt64
		} else {
			res[i] = math.MinInt64
		}
	}
	return res
}

// convert []int to []uint
func IntsToUints(arr []int) []uint {
	res := make([]uint, len(arr))
	for i, v := range arr {
		if v >= 0 {
			res[i] = uint(v)
		} else {
			res[i] = 0
		}
	}
	return res
}

// convert []uint to []int
func UintsToInts(arr []uint) []int {
	res := make([]int, len(arr))
	for i, v := range arr {
		if v <= math.MaxInt64 {
			res[i] = int(v)
		} else {
			res[i] = math.MaxInt64
		}
	}
	return res
}

// convert []uint to []float64
func UintsToFloat64s(arr []uint) []float64 {
	res := make([]float64, len(arr))
	for i, v := range arr {
		res[i] = float64(v)
	}
	return res
}

// convert []float64 to []uint
func Float64sToUints(arr []float64) []uint {
	res := make([]uint, len(arr))
	for i, v := range arr {
		if v <= 0 {
			res[i] = 0
		} else if v <= math.MaxUint64 {
			res[i] = uint(v)
		} else {
			res[i] = math.MaxUint64
		}
	}
	return res
}

// convert []string to []int
func StringsToInts(arr []string) ([]int, error) {
	res := make([]int, len(arr))
	var err error = nil
	var tmp error = nil
	for i, v := range arr {
		res[i], tmp = strconv.Atoi(v)
		if tmp != nil {
			err = tmp
		}
	}
	return res, err
}

// convert []string to []uint
func StringsToUints(arr []string) ([]uint, error) {
	res := make([]uint, len(arr))
	var err error = nil
	var tmp error = nil
	var tmpUint64 uint64
	for i, v := range arr {
		tmpUint64, tmp = strconv.ParseUint(v, 10, 0)
		res[i] = uint(tmpUint64)
		if tmp != nil {
			err = tmp
		}
	}
	return res, err
}

// convert []string to []float64
func StringsToFloat64s(arr []string) ([]float64, error) {
	res := make([]float64, len(arr))
	var err error = nil
	var tmp error = nil
	for i, v := range arr {
		res[i], tmp = strconv.ParseFloat(v, 64)
		if tmp != nil {
			err = tmp
		}
	}
	return res, err
}

// convert []rune to []byte
func Runes2Bytes(rs []rune) []byte {
	size := 0
	for _, r := range rs {
		size += utf8.RuneLen(r)
	}
	bs := make([]byte, size)
	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}
	return bs
}
