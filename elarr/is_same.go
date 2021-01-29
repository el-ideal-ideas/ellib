// This module contains some helper functions for array and slice.

package elarr

// Compare slice data, if all slices are equal, return true. Otherwise, return false.
func IsSameStr(data ...[]string) bool {
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
func IsSameInt(data ...[]int) bool {
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
func IsSameInt8(data ...[]int8) bool {
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
func IsSameInt16(data ...[]int16) bool {
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
func IsSameInt32(data ...[]int32) bool {
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
func IsSameInt64(data ...[]int64) bool {
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
func IsSameUint(data ...[]uint) bool {
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
func IsSameUint8(data ...[]uint8) bool {
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
func IsSameUint16(data ...[]uint16) bool {
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
func IsSameUint32(data ...[]uint32) bool {
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
func IsSameUint64(data ...[]uint64) bool {
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
func IsSameRune(data ...[]rune) bool {
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
func IsSameByte(data ...[]byte) bool {
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
func IsSameFloat32(data ...[]float32) bool {
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
func IsSameFloat64(data ...[]float64) bool {
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
