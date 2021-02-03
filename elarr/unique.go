package elarr

var Unique = UniqueInter

// Remove duplicates item.
// Maybe slow
func UniqueInter(arr []interface{}) []interface{} {
	res := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		if !InInter(res, v) {
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueStr(arr []string) []string {
	res := make([]string, 0, len(arr))
	tmp := make(map[string]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueInt(arr []int) []int {
	res := make([]int, 0, len(arr))
	tmp := make(map[int]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueInt8(arr []int8) []int8 {
	res := make([]int8, 0, len(arr))
	tmp := make(map[int8]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueInt16(arr []int16) []int16 {
	res := make([]int16, 0, len(arr))
	tmp := make(map[int16]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueInt32(arr []int32) []int32 {
	res := make([]int32, 0, len(arr))
	tmp := make(map[int32]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueInt64(arr []int64) []int64 {
	res := make([]int64, 0, len(arr))
	tmp := make(map[int64]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueUint(arr []uint) []uint {
	res := make([]uint, 0, len(arr))
	tmp := make(map[uint]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueUint8(arr []uint8) []uint8 {
	res := make([]uint8, 0, len(arr))
	tmp := make(map[uint8]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueUint16(arr []uint16) []uint16 {
	res := make([]uint16, 0, len(arr))
	tmp := make(map[uint16]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueUint32(arr []uint32) []uint32 {
	res := make([]uint32, 0, len(arr))
	tmp := make(map[uint32]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueUint64(arr []uint64) []uint64 {
	res := make([]uint64, 0, len(arr))
	tmp := make(map[uint64]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueByte(arr []byte) []byte {
	res := make([]byte, 0, len(arr))
	tmp := make(map[byte]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueRune(arr []rune) []rune {
	res := make([]rune, 0, len(arr))
	tmp := make(map[rune]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueFloat32(arr []float32) []float32 {
	res := make([]float32, 0, len(arr))
	tmp := make(map[float32]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Remove duplicates item.
func UniqueFloat64(arr []float64) []float64 {
	res := make([]float64, 0, len(arr))
	tmp := make(map[float64]struct{})
	for _, v := range arr {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}
