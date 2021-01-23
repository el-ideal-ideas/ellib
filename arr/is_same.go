// This module contains some helper functions for array and slice.

package arr

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func StrIsSame(data ...[]string) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IntIsSame(data ...[]int) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Int8IsSame(data ...[]int8) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Int16IsSame(data ...[]int16) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Int32IsSame(data ...[]int32) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Int64IsSame(data ...[]int64) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func UintIsSame(data ...[]uint) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Uint8IsSame(data ...[]uint8) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Uint16IsSame(data ...[]uint16) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Uint32IsSame(data ...[]uint32) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Uint64IsSame(data ...[]uint64) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func RuneIsSame(data ...[]rune) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func ByteIsSame(data ...[]byte) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Float32IsSame(data ...[]float32) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func Float64IsSame(data ...[]float64) bool {
	// check arguments
	count := len(data)
	if count < 2 {
		return false
	}
	// check length
	length := len(data[0])
	for i := 1; i < count; i++ {
		if len(data[i]) != length {
			return false
		}
	}
	// check items
	for i := 0; i < length; i++ {
		first := data[0][i]
		for j := 1; j < count; j++ {
			if first != data[j][i] {
				return false
			}
		}
	}
	return true
}
