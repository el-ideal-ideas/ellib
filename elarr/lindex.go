package elarr

import "reflect"

var LIndex = LIndexInter

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexInter(v []interface{}, item interface{}) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if reflect.DeepEqual(v[i], item) {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexStr(v []string, item string) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexInt(v []int, item int) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexInt8(v []int8, item int8) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexInt16(v []int16, item int16) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexInt32(v []int32, item int32) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexInt64(v []int64, item int64) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexUint(v []uint, item uint) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexUint8(v []uint8, item uint8) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexUint16(v []uint16, item uint16) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexUint32(v []uint32, item uint32) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexUint64(v []uint64, item uint64) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexRune(v []rune, item rune) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexByte(v []byte, item byte) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexFloat32(v []float32, item float32) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the last item that is same as the given `item` parameter.
// If can't found same item, return -1.
func LIndexFloat64(v []float64, item float64) int {
	l := len(v)
	for i := l - 1; i >= 0; i-- {
		if v[i] == item {
			return i
		}
	}
	return -1
}
