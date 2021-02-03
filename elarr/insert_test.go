package elarr

import "testing"

func TestInsert(t *testing.T) {
	d := []interface{}{1, 2}
	d = Insert(d, 0, 0)
	if !IsSame(d, []interface{}{0, 1, 2}) {
		t.Errorf("Invalid result: %v", d)
	}
	d = Insert(d, 1, 9)
	if !IsSame(d, []interface{}{0, 9, 1, 2}) {
		t.Errorf("Invalid result: %v", d)
	}
	d = Insert(d, 3, 3)
	if !IsSame(d, []interface{}{0, 9, 1, 3, 2}) {
		t.Errorf("Invalid result: %v", d)
	}
	d = Insert(d, -5, 0)
	if !IsSame(d, []interface{}{0, 0, 9, 1, 3, 2}) {
		t.Errorf("Invalid result: %v", d)
	}
	d = Insert(d, -5, 1)
	if !IsSame(d, []interface{}{0, 1, 0, 9, 1, 3, 2}) {
		t.Errorf("Invalid result: %v", d)
	}
	d = Insert(d, 7, 7)
	if !IsSame(d, []interface{}{0, 1, 0, 9, 1, 3, 2, 7}) {
		t.Errorf("Invalid result: %v", d)
	}
	ds := []string{"も", "ず", "く"}
	ds = InsertStr(ds, 2, "ず")
	if !IsSame(ds, []string{"も", "ず", "ず", "く"}) {
		t.Errorf("Invalid result: %v", ds)
	}
	di := []int{1, 2, 3}
	di = InsertInt(di, 2, 2)
	if !IsSame(di, []int{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", di)
	}
	di8 := []int8{1, 2, 3}
	di8 = InsertInt8(di8, 2, 2)
	if !IsSame(di8, []int8{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", di8)
	}
	di16 := []int16{1, 2, 3}
	di16 = InsertInt16(di16, 2, 2)
	if !IsSame(di16, []int16{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", di16)
	}
	di32 := []int32{1, 2, 3}
	di32 = InsertInt32(di32, 2, 2)
	if !IsSame(di32, []int32{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", di32)
	}
	di64 := []int64{1, 2, 3}
	di64 = InsertInt64(di64, 2, 2)
	if !IsSame(di64, []int64{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", di64)
	}
	dui := []uint{1, 2, 3}
	dui = InsertUint(dui, 2, 2)
	if !IsSame(dui, []int{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", dui)
	}
	dui8 := []uint8{1, 2, 3}
	dui8 = InsertUint8(dui8, 2, 2)
	if !IsSame(dui8, []int8{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", dui8)
	}
	dui16 := []uint16{1, 2, 3}
	dui16 = InsertUint16(dui16, 2, 2)
	if !IsSame(dui16, []int16{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", dui16)
	}
	dui32 := []uint32{1, 2, 3}
	dui32 = InsertUint32(dui32, 2, 2)
	if !IsSame(dui, []int{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", dui32)
	}
	dui64 := []uint64{1, 2, 3}
	dui64 = InsertUint64(dui64, 2, 2)
	if !IsSame(dui64, []int64{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", dui64)
	}
	r := []rune{1, 2, 3}
	r = InsertRune(r, 2, 2)
	if !IsSame(r, []rune{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", r)
	}
	b := []byte{1, 2, 3}
	b = InsertByte(b, 2, 2)
	if !IsSame(b, []byte{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", b)
	}
	f32 := []float32{1, 2, 3}
	f32 = InsertFloat32(f32, 2, 2)
	if !IsSame(f32, []float32{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", f32)
	}
	f64 := []float64{1, 2, 3}
	f64 = InsertFloat64(f64, 2, 2)
	if !IsSame(f64, []float64{1, 2, 2, 3}) {
		t.Errorf("Invalid result: %v", f64)
	}
	bo := []bool{true, true, true}
	bo = InsertBool(bo, 2, false)
	if !IsSame(bo, []bool{true, true, false, true}) {
		t.Errorf("Invalid result: %v", bo)
	}
}
