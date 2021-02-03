package elarr

var Insert = InsertInter

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertInter(v []interface{}, index int, value interface{}) (newPtr []interface{}) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]interface{}{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertStr(v []string, index int, value string) (newPtr []string) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]string{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertInt(v []int, index int, value int) (newPtr []int) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]int{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertInt8(v []int8, index int, value int8) (newPtr []int8) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]int8{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertInt16(v []int16, index int, value int16) (newPtr []int16) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]int16{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertInt32(v []int32, index int, value int32) (newPtr []int32) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]int32{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertInt64(v []int64, index int, value int64) (newPtr []int64) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]int64{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertUint(v []uint, index int, value uint) (newPtr []uint) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]uint{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertUint8(v []uint8, index int, value uint8) (newPtr []uint8) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]uint8{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertUint16(v []uint16, index int, value uint16) (newPtr []uint16) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]uint16{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertUint32(v []uint32, index int, value uint32) (newPtr []uint32) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]uint32{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertUint64(v []uint64, index int, value uint64) (newPtr []uint64) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]uint64{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertRune(v []rune, index int, value rune) (newPtr []rune) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]rune{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertByte(v []byte, index int, value byte) (newPtr []byte) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]byte{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertFloat32(v []float32, index int, value float32) (newPtr []float32) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]float32{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}

// Insert can insert a value to array&slice at target position.
// index can be negative number.
func InsertFloat64(v []float64, index int, value float64) (newPtr []float64) {
	if index > len(v) || -index > len(v) {
		return v
	}
	if index == len(v) {
		return append(v, value)
	} else if index > 0 {
		newPtr = append(v[:index+1], v[index:]...)
		newPtr[index] = value
	} else if index == 0 {
		newPtr = append([]float64{value}, v...)
	} else if index == -1 {
		newPtr = append(v, value)
	} else {
		newPtr = append(v[:len(v)+index+1], v[len(v)+index:]...)
		newPtr[len(v)+index] = value
	}
	return
}
