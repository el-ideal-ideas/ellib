package elarr

var LRemove = LRemoveInter

// LRemove the last element that is same as the given item.
func LRemoveInter(v []interface{}, item interface{}) (newPtr []interface{}) {
	_, newPtr = PopInter(v, LIndexInter(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveStr(v []string, item string) (newPtr []string) {
	_, newPtr = PopStr(v, LIndexStr(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveInt(v []int, item int) (newPtr []int) {
	_, newPtr = PopInt(v, LIndexInt(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveInt8(v []int8, item int8) (newPtr []int8) {
	_, newPtr = PopInt8(v, LIndexInt8(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveInt16(v []int16, item int16) (newPtr []int16) {
	_, newPtr = PopInt16(v, LIndexInt16(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveInt32(v []int32, item int32) (newPtr []int32) {
	_, newPtr = PopInt32(v, LIndexInt32(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveInt64(v []int64, item int64) (newPtr []int64) {
	_, newPtr = PopInt64(v, LIndexInt64(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveUint(v []uint, item uint) (newPtr []uint) {
	_, newPtr = PopUint(v, LIndexUint(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveUint8(v []uint8, item uint8) (newPtr []uint8) {
	_, newPtr = PopUint8(v, LIndexUint8(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveUint16(v []uint16, item uint16) (newPtr []uint16) {
	_, newPtr = PopUint16(v, LIndexUint16(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveUint32(v []uint32, item uint32) (newPtr []uint32) {
	_, newPtr = PopUint32(v, LIndexUint32(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveUint64(v []uint64, item uint64) (newPtr []uint64) {
	_, newPtr = PopUint64(v, LIndexUint64(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveRune(v []rune, item rune) (newPtr []rune) {
	_, newPtr = PopRune(v, LIndexRune(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveByte(v []byte, item byte) (newPtr []byte) {
	_, newPtr = PopByte(v, LIndexByte(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveFloat32(v []float32, item float32) (newPtr []float32) {
	_, newPtr = PopFloat32(v, LIndexFloat32(v, item))
	return
}

// LRemove the last element that is same as the given item.
func LRemoveFloat64(v []float64, item float64) (newPtr []float64) {
	_, newPtr = PopFloat64(v, LIndexFloat64(v, item))
	return
}
