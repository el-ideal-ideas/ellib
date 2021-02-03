package elarr

var Remove = RemoveInter

// Remove the first element that is same as the given item.
func RemoveInter(v []interface{}, item interface{}) (newPtr []interface{}) {
	_, newPtr = PopInter(v, IndexInter(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveStr(v []string, item string) (newPtr []string) {
	_, newPtr = PopStr(v, IndexStr(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveInt(v []int, item int) (newPtr []int) {
	_, newPtr = PopInt(v, IndexInt(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveInt8(v []int8, item int8) (newPtr []int8) {
	_, newPtr = PopInt8(v, IndexInt8(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveInt16(v []int16, item int16) (newPtr []int16) {
	_, newPtr = PopInt16(v, IndexInt16(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveInt32(v []int32, item int32) (newPtr []int32) {
	_, newPtr = PopInt32(v, IndexInt32(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveInt64(v []int64, item int64) (newPtr []int64) {
	_, newPtr = PopInt64(v, IndexInt64(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveUint(v []uint, item uint) (newPtr []uint) {
	_, newPtr = PopUint(v, IndexUint(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveUint8(v []uint8, item uint8) (newPtr []uint8) {
	_, newPtr = PopUint8(v, IndexUint8(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveUint16(v []uint16, item uint16) (newPtr []uint16) {
	_, newPtr = PopUint16(v, IndexUint16(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveUint32(v []uint32, item uint32) (newPtr []uint32) {
	_, newPtr = PopUint32(v, IndexUint32(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveUint64(v []uint64, item uint64) (newPtr []uint64) {
	_, newPtr = PopUint64(v, IndexUint64(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveRune(v []rune, item rune) (newPtr []rune) {
	_, newPtr = PopRune(v, IndexRune(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveByte(v []byte, item byte) (newPtr []byte) {
	_, newPtr = PopByte(v, IndexByte(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveFloat32(v []float32, item float32) (newPtr []float32) {
	_, newPtr = PopFloat32(v, IndexFloat32(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveFloat64(v []float64, item float64) (newPtr []float64) {
	_, newPtr = PopFloat64(v, IndexFloat64(v, item))
	return
}

// Remove the first element that is same as the given item.
func RemoveBool(v []bool, item bool) (newPtr []bool) {
	_, newPtr = PopBool(v, IndexBool(v, item))
	return
}
