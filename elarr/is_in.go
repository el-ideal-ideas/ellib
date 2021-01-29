// This module contains some helper functions for array and slice.
// For example:
// var s = []string{"も", "ず", "く"}
// fmt.Println(InStr(s, "も"))
// >> true

package elarr

// If `item` in `data` return true else return false.
// call function `f` for all items in `data`, When `f` returns true,
// exit from the loop, and return true.
func InF(data []interface{}, f func(interface{}) bool) bool {
	for _, i := range data {
		if f(i) {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InStr(data []string, item string) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InInt(data []int, item int) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InInt8(data []int8, item int8) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InInt16(data []int16, item int16) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InInt32(data []int32, item int32) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InInt64(data []int64, item int64) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InUint(data []uint, item uint) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InUint8(data []uint8, item uint8) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InUint16(data []uint16, item uint16) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InUint32(data []uint32, item uint32) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InUint64(data []uint64, item uint64) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InRune(data []rune, item rune) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InByte(data []byte, item byte) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InFloat32(data []float32, item float32) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}

// If `item` in `data` return true else return false.
func InFloat64(data []float64, item float64) bool {
	for _, i := range data {
		if i == item {
			return true
		}
	}
	return false
}
