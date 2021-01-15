// This module contains some helper functions for array and slice.
// For example:
// var s = []string{"も", "ず", "く"}
// fmt.Println(StrIn(s, "も"))
// >> true

package arr


// If `item` in `data` return true else return false.
func StrIn(data []string, item string) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func IntIn(data []int, item int) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Int8In(data []int8, item int8) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Int16In(data []int16, item int16) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Int32In(data []int32, item int32) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Int64In(data []int64, item int64) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func UintIn(data []uint, item uint) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Uint8In(data []uint8, item uint8) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Uint16In(data []uint16, item uint16) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Uint32In(data []uint32, item uint32) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Uint64In(data []uint64, item uint64) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func RuneIn(data []rune, item rune) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func ByteIn(data []byte, item byte) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Float32In(data []float32, item float32) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func Float64In(data []float64, item float64) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}