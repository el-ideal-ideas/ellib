package elarr

import "testing"

func TestIsSame(t *testing.T) {
	if !IsSameInter([]interface{}{"も", "ず", "く"}, []string{"も", "ず", "く"}, []string{"も", "ず", "く"}) {
		t.Errorf("Invalid result")
	}
	if !IsSameInterSlice([]interface{}{"も", "ず", "く"}, []interface{}{"も", "ず", "く"}, []interface{}{"も", "ず", "く"}) {
		t.Errorf("Invalid result")
	}
	if !IsSameStr([]string{"も", "ず", "く"}, []string{"も", "ず", "く"}, []string{"も", "ず", "く"}) {
		t.Errorf("Invalid result")
	}
	if IsSameStr([]string{"あ", "ず", "く"}, []string{"も", "ず", "く"}, []string{"も", "ず", "く"}) {
		t.Errorf("Invalid result")
	}
	if !IsSameInt([]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameInt([]int{0, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameInt8([]int8{1, 2, 3}, []int8{1, 2, 3}, []int8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameInt8([]int8{0, 2, 3}, []int8{1, 2, 3}, []int8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameInt16([]int16{1, 2, 3}, []int16{1, 2, 3}, []int16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameInt16([]int16{0, 2, 3}, []int16{1, 2, 3}, []int16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameInt32([]int32{1, 2, 3}, []int32{1, 2, 3}, []int32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameInt32([]int32{0, 2, 3}, []int32{1, 2, 3}, []int32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameInt64([]int64{1, 2, 3}, []int64{1, 2, 3}, []int64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameInt64([]int64{0, 2, 3}, []int64{1, 2, 3}, []int64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameUint([]uint{1, 2, 3}, []uint{1, 2, 3}, []uint{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameUint([]uint{0, 2, 3}, []uint{1, 2, 3}, []uint{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameUint8([]uint8{1, 2, 3}, []uint8{1, 2, 3}, []uint8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameUint8([]uint8{0, 2, 3}, []uint8{1, 2, 3}, []uint8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameUint16([]uint16{1, 2, 3}, []uint16{1, 2, 3}, []uint16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameUint16([]uint16{0, 2, 3}, []uint16{1, 2, 3}, []uint16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameUint32([]uint32{1, 2, 3}, []uint32{1, 2, 3}, []uint32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameUint32([]uint32{0, 2, 3}, []uint32{1, 2, 3}, []uint32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameUint64([]uint64{1, 2, 3}, []uint64{1, 2, 3}, []uint64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameUint64([]uint64{0, 2, 3}, []uint64{1, 2, 3}, []uint64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameRune([]rune{'も', 'ず', 'く'}, []rune{'も', 'ず', 'く'}, []rune{'も', 'ず', 'く'}) {
		t.Errorf("Invalid result")
	}
	if IsSameRune([]rune{'あ', 'ず', 'く'}, []rune{'も', 'ず', 'く'}, []rune{'も', 'ず', 'く'}) {
		t.Errorf("Invalid result")
	}
	if !IsSameByte([]byte{1, 2, 3}, []byte{1, 2, 3}, []byte{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IsSameByte([]byte{0, 2, 3}, []byte{1, 2, 3}, []byte{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsSameFloat32([]float32{1.0, 2.0, 3.0}, []float32{1.0, 2.0, 3.0}, []float32{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
	if IsSameFloat32([]float32{0.0, 2.0, 3.0}, []float32{1.0, 2.0, 3.0}, []float32{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
	if !IsSameFloat64([]float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
	if IsSameFloat64([]float64{0.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
}
