package elarr

import "testing"

func TestReverse(t *testing.T) {
	d1 := []string{"も", "ず", "く"}
	ReverseStr(d1)
	if !IsSameStr(d1, []string{"く", "ず", "も"}) {
		t.Errorf("Invalid result")
	}
	d2 := []int{1, 2, 3}
	ReverseInt(d2)
	if !IsSameInt(d2, []int{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d3 := []int8{1, 2, 3}
	ReverseInt8(d3)
	if !IsSameInt8(d3, []int8{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d4 := []int16{1, 2, 3}
	ReverseInt16(d4)
	if !IsSameInt16(d4, []int16{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d5 := []int32{1, 2, 3}
	ReverseInt32(d5)
	if !IsSameInt32(d5, []int32{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d6 := []int64{1, 2, 3}
	ReverseInt64(d6)
	if !IsSameInt64(d6, []int64{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d7 := []uint{1, 2, 3}
	ReverseUint(d7)
	if !IsSameUint(d7, []uint{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d8 := []uint8{1, 2, 3}
	ReverseUint8(d8)
	if !IsSameUint8(d8, []uint8{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d9 := []uint16{1, 2, 3}
	ReverseUint16(d9)
	if !IsSameUint16(d9, []uint16{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d10 := []uint32{1, 2, 3}
	ReverseUint32(d10)
	if !IsSameUint32(d10, []uint32{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d11 := []uint64{1, 2, 3}
	ReverseUint64(d11)
	if !IsSameUint64(d11, []uint64{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d12 := []rune{'も', 'ず', 'く'}
	ReverseRune(d12)
	if !IsSameRune(d12, []rune{'く', 'ず', 'も'}) {
		t.Errorf("Invalid result")
	}
	d13 := []byte{1, 2, 3}
	ReverseByte(d13)
	if !IsSameByte(d13, []byte{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d14 := []float32{1.0, 2.0, 3.0}
	ReverseFloat32(d14)
	if !IsSameFloat32(d14, []float32{3.0, 2.0, 1.0}) {
		t.Errorf("Invalid result")
	}
	d15 := []float64{1.0, 2.0, 3.0}
	ReverseFloat64(d15)
	if !IsSameFloat64(d15, []float64{3.0, 2.0, 1.0}) {
		t.Errorf("Invalid result")
	}
	d16 := []interface{}{1.0, 2.0, 3.0}
	Reverse(d16)
	if !IsSameFloat64(ToFloat64s(d16), []float64{3.0, 2.0, 1.0}) {
		t.Errorf("Invalid result")
	}
}
