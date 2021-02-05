## elmath
- 数学方面的工具包
- utilities for math
- 数学系のユーティリティ
```go
import "github.com/el-ideal-ideas/ellib/elmath"

// 四捨五入
// Round half away from zero
// precision can be a negative number.
// 四舍五入
elmath.Round(val float64, precision int) float64

// 计算百分比
// get values percent of the total
// パーセントの値に変換する
elmath.Percent(val, total int) float64

// 阶乘
// Factorial
// 階乗
// 5! == 5*4*3*2*1 == 120
elmath.Factorial(n uint) uint

// 排列
// Permutation
// 順列
// Permutation(5, 3) == 5P3 == 5*4*3
elmath.Permutation(a, b uint) uint

// 组合
// Combination
// 組み合わせ
// Combination(5, 3) == 5C3 == (5*4*3) / (3*2*1) == 10
elmath.Combination(n uint, k uint) uint

```