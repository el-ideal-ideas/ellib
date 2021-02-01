package elmath

import "math"

// Round half away from zero
// precision can be a negative number.
// 四捨五入
// 四舍五入
func Round(val float64, precision int) float64 {
	if precision == 0 {
		return math.Round(val)
	}
	p := math.Pow10(precision)
	if precision < 0 {
		return math.Floor(val*p+0.5) * math.Pow10(-precision)
	}
	return math.Floor(val*p+0.5) / p
}

// Percent returns a values percent of the total
func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}
	return (float64(val) / float64(total)) * 100
}

// Factorial
// 階乗
// 5! == 5*4*3*2*1 == 120
func Factorial(n uint) uint {
	return Permutation(n, n-1)
}

// Permutation
// 順列
// Permutation(5, 3) == 5P3 == 5*4*3
func Permutation(a, b uint) uint {
	if b > a {
		return 0
	} else if b > 0 && a >= b {
		var i uint
		var res uint = 1
		for i = 0; i < b; i++ {
			res *= a
			a -= 1
		}
		return res
	} else {
		return 1
	}
}

// Combination
// 組み合わせ
// Combination(5, 3) == 5C3 == (5*4*3) / (3*2*1) == 10
func Combination(n uint, k uint) uint {
	return Permutation(n, k) / Factorial(k)
}
