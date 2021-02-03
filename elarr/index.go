package elarr

import "reflect"

var Index = IndexInter

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexInter(v []interface{}, item interface{}) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if reflect.DeepEqual(v[i], item) {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexStr(v []string, item string) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexInt(v []int, item int) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexInt8(v []int8, item int8) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexInt16(v []int16, item int16) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexInt32(v []int32, item int32) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexInt64(v []int64, item int64) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexUint(v []uint, item uint) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexUint8(v []uint8, item uint8) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexUint16(v []uint16, item uint16) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexUint32(v []uint32, item uint32) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexUint64(v []uint64, item uint64) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexRune(v []rune, item rune) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexByte(v []byte, item byte) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexFloat32(v []float32, item float32) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexFloat64(v []float64, item float64) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}

// Get the index of the first item that is same as the given `item` parameter.
// If can't found same item, return -1.
func IndexBool(v []bool, item bool) int {
	l := len(v)
	for i := 0; i < l; i++ {
		if v[i] == item {
			return i
		}
	}
	return -1
}
