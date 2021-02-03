package elarr

import "testing"

func TestLRemove(t *testing.T) {
	d := []interface{}{1, 3, 2, 3, 4}
	d = LRemoveInter(d, 3)
	if !IsSame(d, []int{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	ds := []string{"1", "3", "2", "3", "4"}
	ds = LRemoveStr(ds, "3")
	if !IsSame(d, []string{"1", "3", "2", "4"}) {
		t.Errorf("Invalid result")
	}
	di := []int{1, 3, 2, 3, 4}
	di = LRemoveInt(di, 3)
	if !IsSame(di, []int{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	di8 := []int8{1, 3, 2, 3, 4}
	di8 = LRemoveInt8(di8, 3)
	if !IsSame(di8, []int8{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	di16 := []int16{1, 3, 2, 3, 4}
	di16 = LRemoveInt16(di16, 3)
	if !IsSame(di16, []int16{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	di32 := []int32{1, 3, 2, 3, 4}
	di32 = LRemoveInt32(di32, 3)
	if !IsSame(di32, []int32{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	di64 := []int64{1, 3, 2, 3, 4}
	di64 = LRemoveInt64(di64, 3)
	if !IsSame(di64, []int64{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui := []uint{1, 3, 2, 3, 4}
	dui = LRemoveUint(dui, 3)
	if !IsSame(dui, []uint{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui8 := []uint8{1, 3, 2, 3, 4}
	dui8 = LRemoveUint8(dui8, 3)
	if !IsSame(dui8, []uint8{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui16 := []uint16{1, 3, 2, 3, 4}
	dui16 = LRemoveUint16(dui16, 3)
	if !IsSame(dui16, []uint16{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui32 := []uint32{1, 3, 2, 3, 4}
	dui32 = LRemoveUint32(dui32, 3)
	if !IsSame(dui32, []uint32{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui64 := []uint64{1, 3, 2, 3, 4}
	dui64 = LRemoveUint64(dui64, 3)
	if !IsSame(dui64, []uint64{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	r := []rune{1, 3, 2, 3, 4}
	r = LRemoveRune(r, 3)
	if !IsSame(r, []rune{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	b := []byte{1, 3, 2, 3, 4}
	b = LRemoveByte(b, 3)
	if !IsSame(b, []byte{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	f32 := []float32{1, 3, 2, 3, 4}
	f32 = LRemoveFloat32(f32, 3)
	if !IsSame(f32, []float32{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
	f64 := []float64{1, 3, 2, 3, 4}
	f64 = LRemoveFloat64(f64, 3)
	if !IsSame(f64, []float64{1, 3, 2, 4}) {
		t.Errorf("Invalid result")
	}
}
