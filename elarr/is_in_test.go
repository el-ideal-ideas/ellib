package elarr

import "testing"

func TestIsIn(t *testing.T) {
	if !InF([]interface{}{"も", "ず", "く"}, func(v interface{}) bool { return v.(string) == "く" }) {
		t.Errorf("Invalid result")
	}
	if !InInter([]interface{}{"も", "ず", "く"}, "も") {
		t.Errorf("Invalid result")
	}
	if !InStr([]string{"も", "ず", "く"}, "も") {
		t.Errorf("Invalid result")
	}
	if InStr([]string{"も", "ず", "く"}, "あ") {
		t.Errorf("Invalid result")
	}
	if !InInt([]int{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InInt([]int{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InInt8([]int8{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InInt8([]int8{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InInt16([]int16{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InInt16([]int16{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InInt32([]int32{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InInt32([]int32{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InInt64([]int64{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InInt64([]int64{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InUint([]uint{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InUint([]uint{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InUint8([]uint8{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InUint8([]uint8{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InUint16([]uint16{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InUint16([]uint16{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InUint32([]uint32{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InUint32([]uint32{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InUint64([]uint64{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InUint64([]uint64{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InRune([]rune{'も', 'ず', 'く'}, 'も') {
		t.Errorf("Invalid result")
	}
	if InRune([]rune{'も', 'ず', 'く'}, 'あ') {
		t.Errorf("Invalid result")
	}
	if !InByte([]byte{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if InByte([]byte{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !InFloat32([]float32{1.0, 2.0, 3.0}, 1.0) {
		t.Errorf("Invalid result")
	}
	if InFloat32([]float32{1.0, 2.0, 3.0}, 0.0) {
		t.Errorf("Invalid result")
	}
	if !InFloat64([]float64{1.0, 2.0, 3.0}, 1.0) {
		t.Errorf("Invalid result")
	}
	if InFloat64([]float64{1.0, 2.0, 3.0}, 0.0) {
		t.Errorf("Invalid result")
	}
	if InBool([]bool{true, true}, false) {
		t.Errorf("Invalid result")
	}
}
