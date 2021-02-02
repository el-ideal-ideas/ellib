package elarr

import (
	"testing"
)

func TestFill(t *testing.T) {
	d0 := make([]interface{}, 3)
	FillInter(d0, "el")
	if !IsSameInterSlice(d0, []interface{}{"el", "el", "el"}) {
		t.Errorf("Invalid result")
	}
	d1 := make([]string, 3)
	FillStr(d1, "el")
	if !IsSameStr(d1, []string{"el", "el", "el"}) {
		t.Errorf("Invalid result")
	}
	d2 := make([]int, 3)
	FillInt(d2, 1)
	if !IsSameInt(d2, []int{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d3 := make([]int8, 3)
	FillInt8(d3, 1)
	if !IsSameInt8(d3, []int8{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d4 := make([]int16, 3)
	FillInt16(d4, 1)
	if !IsSameInt16(d4, []int16{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d5 := make([]int32, 3)
	FillInt32(d5, 1)
	if !IsSameInt32(d5, []int32{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d6 := make([]int64, 3)
	FillInt64(d6, 1)
	if !IsSameInt64(d6, []int64{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d7 := make([]uint, 3)
	FillUint(d7, 1)
	if !IsSameUint(d7, []uint{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d8 := make([]uint8, 3)
	FillUint8(d8, 1)
	if !IsSameUint8(d8, []uint8{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d9 := make([]uint16, 3)
	FillUint16(d9, 1)
	if !IsSameUint16(d9, []uint16{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d10 := make([]uint32, 3)
	FillUint32(d10, 1)
	if !IsSameUint32(d10, []uint32{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d11 := make([]uint64, 3)
	FillUint64(d11, 1)
	if !IsSameUint64(d11, []uint64{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d12 := make([]rune, 3)
	FillRune(d12, 1)
	if !IsSameRune(d12, []rune{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d13 := make([]byte, 3)
	FillByte(d13, 1)
	if !IsSameByte(d13, []byte{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d14 := make([]float32, 3)
	FillFloat32(d14, 1)
	if !IsSameFloat32(d14, []float32{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
	d15 := make([]float64, 3)
	FillFloat64(d15, 1)
	if !IsSameFloat64(d15, []float64{1, 1, 1}) {
		t.Errorf("Invalid result")
	}
}
