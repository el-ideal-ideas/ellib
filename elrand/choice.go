package elrand

import "reflect"

// Choice some random item from data.
// supports array, slice, map.
// This function use reflect, will be more slow than other Choice functions.
func Choice(data interface{}, count int) []interface{} {
	if count < 0 {
		return []interface{}{}
	}
	v := reflect.ValueOf(data)
	res := make([]interface{}, 0, count)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < count; i++ {
			res = append(res, v.Index(RandomInt(0, v.Len())).Interface())
		}
	case reflect.Map:
		keys := v.MapKeys()
		for i := 0; i < count; i++ {
			res = append(res, v.MapIndex(keys[RandomInt(0, len(keys))]).Interface())
		}
	}
	return res
}

// Choice some random item from data.
func ChoiceInter(data []interface{}, count int) []interface{} {
	if count < 0 {
		return []interface{}{}
	}
	res := make([]interface{}, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceStr(data []string, count int) []string {
	if count < 0 {
		return []string{}
	}
	res := make([]string, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceInt(data []int, count int) []int {
	if count < 0 {
		return []int{}
	}
	res := make([]int, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceInt8(data []int8, count int) []int8 {
	if count < 0 {
		return []int8{}
	}
	res := make([]int8, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceInt16(data []int16, count int) []int16 {
	if count < 0 {
		return []int16{}
	}
	res := make([]int16, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceInt32(data []int32, count int) []int32 {
	if count < 0 {
		return []int32{}
	}
	res := make([]int32, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceInt64(data []int64, count int) []int64 {
	if count < 0 {
		return []int64{}
	}
	res := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceUint(data []uint, count int) []uint {
	if count < 0 {
		return []uint{}
	}
	res := make([]uint, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceUint8(data []uint8, count int) []uint8 {
	if count < 0 {
		return []uint8{}
	}
	res := make([]uint8, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceUint16(data []uint16, count int) []uint16 {
	if count < 0 {
		return []uint16{}
	}
	res := make([]uint16, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceUint32(data []uint32, count int) []uint32 {
	if count < 0 {
		return []uint32{}
	}
	res := make([]uint32, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceUint64(data []uint64, count int) []uint64 {
	if count < 0 {
		return []uint64{}
	}
	res := make([]uint64, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceRune(data []rune, count int) []rune {
	if count < 0 {
		return []rune{}
	}
	res := make([]rune, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceByte(data []byte, count int) []byte {
	if count < 0 {
		return []byte{}
	}
	res := make([]byte, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceFloat32(data []float32, count int) []float32 {
	if count < 0 {
		return []float32{}
	}
	res := make([]float32, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}

// Choice some random item from data.
func ChoiceFloat64(data []float64, count int) []float64 {
	if count < 0 {
		return []float64{}
	}
	res := make([]float64, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, data[RandomInt(0, len(data))])
	}
	return res
}
