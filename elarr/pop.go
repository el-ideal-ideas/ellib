// Usage:
// d := []interface{}{1, 2, 3, 4}
// var v int
// v, d = Pop(d, 2)

package elarr

var Pop = PopInter

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopInter(v []interface{}, index int) (res interface{}, newPtr []interface{}) {
	if index >= len(v) || -index > len(v) {
		return nil, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopStr(v []string, index int) (res string, newPtr []string) {
	if index >= len(v) || -index > len(v) {
		return "", v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopInt(v []int, index int) (res int, newPtr []int) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopInt8(v []int8, index int) (res int8, newPtr []int8) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopInt16(v []int16, index int) (res int16, newPtr []int16) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopInt32(v []int32, index int) (res int32, newPtr []int32) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopInt64(v []int64, index int) (res int64, newPtr []int64) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopUint(v []uint, index int) (res uint, newPtr []uint) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopUint8(v []uint8, index int) (res uint8, newPtr []uint8) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopUint16(v []uint16, index int) (res uint16, newPtr []uint16) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopUint32(v []uint32, index int) (res uint32, newPtr []uint32) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopUint64(v []uint64, index int) (res uint64, newPtr []uint64) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopRune(v []rune, index int) (res rune, newPtr []rune) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopByte(v []byte, index int) (res byte, newPtr []byte) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopFloat32(v []float32, index int) (res float32, newPtr []float32) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopFloat64(v []float64, index int) (res float64, newPtr []float64) {
	if index >= len(v) || -index > len(v) {
		return 0, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}

// Pop can get one item and remove it from array&slice.
// index can be negative number.
func PopBool(v []bool, index int) (res bool, newPtr []bool) {
	if index >= len(v) || -index > len(v) {
		return false, v
	}
	if index > 0 {
		res = v[index]
		newPtr = append(v[:index], v[index+1:]...)
	} else if index == 0 {
		res = v[0]
		newPtr = v[1:]
	} else if index == -1 {
		res = v[len(v)-1]
		newPtr = v[:len(v)-1]
	} else {
		res = v[len(v)+index]
		newPtr = append(v[:len(v)+index], v[len(v)+index+1:]...)
	}
	return
}
