package elarr

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapStrInter(arr []string, f func(string) interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapStrStr(arr []string, f func(string) string) []string {
	res := make([]string, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapStrInt(arr []string, f func(string) int) []int {
	res := make([]int, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapStrUint(arr []string, f func(string) uint) []uint {
	res := make([]uint, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapStrFloat64(arr []string, f func(string) float64) []float64 {
	res := make([]float64, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapIntInter(arr []int, f func(int) interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapIntStr(arr []int, f func(int) string) []string {
	res := make([]string, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapIntInt(arr []int, f func(int) int) []int {
	res := make([]int, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapIntUint(arr []int, f func(int) uint) []uint {
	res := make([]uint, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapFloat64Inter(arr []float64, f func(float64) interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapIntFloat64(arr []int, f func(int) float64) []float64 {
	res := make([]float64, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapFloat64Str(arr []float64, f func(float64) string) []string {
	res := make([]string, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapFloat64Int(arr []float64, f func(float64) int) []int {
	res := make([]int, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapFloat64Uint(arr []float64, f func(float64) uint) []uint {
	res := make([]uint, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapFloat64Float64(arr []float64, f func(float64) float64) []float64 {
	res := make([]float64, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapUintInter(arr []uint, f func(uint) interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapUintStr(arr []uint, f func(uint) string) []string {
	res := make([]string, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapUintInt(arr []uint, f func(uint) int) []int {
	res := make([]int, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapUintUint(arr []uint, f func(uint) uint) []uint {
	res := make([]uint, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapUintFloat64(arr []uint, f func(uint) float64) []float64 {
	res := make([]float64, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapInterInter(arr []interface{}, f func(interface{}) interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapInterStr(arr []interface{}, f func(interface{}) string) []string {
	res := make([]string, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapInterInt(arr []interface{}, f func(interface{}) int) []int {
	res := make([]int, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapInterUint(arr []interface{}, f func(interface{}) uint) []uint {
	res := make([]uint, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}

// Call `f` for all item in `arr`, get the response as a slice.
// Like python's map function.
func MapInterFloat64(arr []interface{}, f func(interface{}) float64) []float64 {
	res := make([]float64, len(arr))
	l := len(arr)
	for i := 0; i < l; i++ {
		res[0] = f(arr[0])
	}
	return res
}
