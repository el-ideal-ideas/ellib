package elarr

var ForEach = ForEachInter

// Call `f` for all item in `arr`
func ForEachInter(arr []interface{}, f func(interface{})) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachStr(arr []string, f func(string)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachInt(arr []int, f func(int)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachInt8(arr []int8, f func(int8)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachInt16(arr []int16, f func(int16)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachInt32(arr []int32, f func(int32)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachInt64(arr []int64, f func(int64)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachUint(arr []uint, f func(uint)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachUint8(arr []uint8, f func(uint8)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachUint16(arr []uint16, f func(uint16)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachUint32(arr []uint32, f func(uint32)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachUint64(arr []uint64, f func(uint64)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachRune(arr []rune, f func(rune)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachByte(arr []byte, f func(byte)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachFloat32(arr []float32, f func(float32)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachFloat64(arr []float64, f func(float64)) {
	for i := range arr {
		f(arr[i])
	}
}

// Call `f` for all item in `arr`
func ForEachBool(arr []bool, f func(bool)) {
	for i := range arr {
		f(arr[i])
	}
}
