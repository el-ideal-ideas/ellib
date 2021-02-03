package elarr

import "testing"

func TestRemove(t *testing.T) {
	d := []interface{}{1, 2, 3, 4}
	d = RemoveInter(d, 3)
	if !IsSame(d, []int{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	ds := []string{"1", "2", "3", "4"}
	ds = RemoveStr(ds, "3")
	if !IsSame(d, []string{"1", "2", "4"}) {
		t.Errorf("Invalid result")
	}
	di := []int{1, 2, 3, 4}
	di = RemoveInt(di, 3)
	if !IsSame(di, []int{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	di8 := []int8{1, 2, 3, 4}
	di8 = RemoveInt8(di8, 3)
	if !IsSame(di8, []int8{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	di16 := []int16{1, 2, 3, 4}
	di16 = RemoveInt16(di16, 3)
	if !IsSame(di16, []int16{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	di32 := []int32{1, 2, 3, 4}
	di32 = RemoveInt32(di32, 3)
	if !IsSame(di32, []int32{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	di64 := []int64{1, 2, 3, 4}
	di64 = RemoveInt64(di64, 3)
	if !IsSame(di64, []int64{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui := []uint{1, 2, 3, 4}
	dui = RemoveUint(dui, 3)
	if !IsSame(dui, []uint{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui8 := []uint8{1, 2, 3, 4}
	dui8 = RemoveUint8(dui8, 3)
	if !IsSame(dui8, []uint8{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui16 := []uint16{1, 2, 3, 4}
	dui16 = RemoveUint16(dui16, 3)
	if !IsSame(dui16, []uint16{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui32 := []uint32{1, 2, 3, 4}
	dui32 = RemoveUint32(dui32, 3)
	if !IsSame(dui32, []uint32{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	dui64 := []uint64{1, 2, 3, 4}
	dui64 = RemoveUint64(dui64, 3)
	if !IsSame(dui64, []uint64{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	r := []rune{1, 2, 3, 4}
	r = RemoveRune(r, 3)
	if !IsSame(r, []rune{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	b := []byte{1, 2, 3, 4}
	b = RemoveByte(b, 3)
	if !IsSame(b, []byte{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	f32 := []float32{1, 2, 3, 4}
	f32 = RemoveFloat32(f32, 3)
	if !IsSame(f32, []float32{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	f64 := []float64{1, 2, 3, 4}
	f64 = RemoveFloat64(f64, 3)
	if !IsSame(f64, []float64{1, 2, 4}) {
		t.Errorf("Invalid result")
	}
	bo := []bool{true, true, false, true}
	bo = LRemoveBool(bo, false)
	if !IsSame(bo, []bool{true, true, true}) {
		t.Errorf("Invalid result")
	}
}
