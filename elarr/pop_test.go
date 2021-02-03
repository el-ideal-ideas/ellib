package elarr

import "testing"

func TestPop(t *testing.T) {
	d := []interface{}{1, 2, 3, 4, 5}
	v, d := Pop(d, 0)
	if v != 1 || !IsSame(d, []interface{}{2, 3, 4, 5}) {
		t.Errorf("Invalid result: %v, %v", v, d)
	}
	d = []interface{}{1, 2, 3, 4, 5}
	v, d = Pop(d, 3)
	if v != 4 || !IsSame(d, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", v, d)
	}
	d = []interface{}{1, 2, 3, 4, 5}
	v, d = Pop(d, 4)
	if v != 5 || !IsSame(d, []interface{}{1, 2, 3, 4}) {
		t.Errorf("Invalid result: %v, %v", v, d)
	}
	d = []interface{}{1, 2, 3, 4, 5}
	v, d = Pop(d, -1)
	if v != 5 || !IsSame(d, []interface{}{1, 2, 3, 4}) {
		t.Errorf("Invalid result: %v, %v", v, d)
	}
	d = []interface{}{1, 2, 3, 4, 5}
	v, d = Pop(d, -2)
	if v != 4 || !IsSame(d, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", v, d)
	}
	d = []interface{}{1, 2, 3, 4, 5}
	v, d = Pop(d, -5)
	if v != 1 || !IsSame(d, []interface{}{2, 3, 4, 5}) {
		t.Errorf("Invalid result: %v, %v", v, d)
	}
	di := []int{1, 2, 3, 4, 5}
	vi, di := PopInt(di, 3)
	if vi != 4 || !IsSame(di, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vi, di)
	}
	di8 := []int8{1, 2, 3, 4, 5}
	vi8, di8 := PopInt8(di8, 3)
	if vi8 != 4 || !IsSame(di8, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vi8, di8)
	}
	di16 := []int16{1, 2, 3, 4, 5}
	vi16, di16 := PopInt16(di16, 3)
	if vi16 != 4 || !IsSame(di16, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vi16, di16)
	}
	di32 := []int32{1, 2, 3, 4, 5}
	vi32, di32 := PopInt32(di32, 3)
	if vi32 != 4 || !IsSame(di32, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vi32, di32)
	}
	di64 := []int64{1, 2, 3, 4, 5}
	vi64, di64 := PopInt64(di64, 3)
	if vi != 4 || !IsSame(di, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vi64, di64)
	}
	dui := []uint{1, 2, 3, 4, 5}
	vui, dui := PopUint(dui, 3)
	if vi != 4 || !IsSame(dui, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vui, dui)
	}
	dui8 := []uint8{1, 2, 3, 4, 5}
	vui8, dui8 := PopUint8(dui8, 3)
	if vi8 != 4 || !IsSame(dui8, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vui8, dui8)
	}
	dui16 := []uint16{1, 2, 3, 4, 5}
	vui16, dui16 := PopUint16(dui16, 3)
	if vi16 != 4 || !IsSame(dui16, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vui16, dui16)
	}
	dui32 := []uint32{1, 2, 3, 4, 5}
	vui32, dui32 := PopUint32(dui32, 3)
	if vi32 != 4 || !IsSame(dui32, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vui32, dui32)
	}
	dui64 := []uint64{1, 2, 3, 4, 5}
	vui64, dui64 := PopUint64(dui64, 3)
	if vi64 != 4 || !IsSame(dui64, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vui64, dui64)
	}
	dr := []rune{1, 2, 3, 4, 5}
	vr, dr := PopRune(dr, 3)
	if vr != 4 || !IsSame(dr, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vr, dr)
	}
	db := []byte{1, 2, 3, 4, 5}
	vb, db := PopByte(db, 3)
	if vb != 4 || !IsSame(db, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vb, db)
	}
	df32 := []float32{1, 2, 3, 4, 5}
	vf32, df32 := PopFloat32(df32, 3)
	if vf32 != 4 || !IsSame(df32, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vf32, df32)
	}
	df64 := []float64{1, 2, 3, 4, 5}
	vf64, df64 := PopFloat64(df64, 3)
	if vf64 != 4 || !IsSame(df64, []interface{}{1, 2, 3, 5}) {
		t.Errorf("Invalid result: %v, %v", vf64, df64)
	}
	ds := []string{"も", "ず", "く"}
	vs, ds := PopStr(ds, 1)
	if vs != "ず" || !IsSame(ds, []interface{}{"も", "く"}) {
		t.Errorf("Invalid result: %v, %v", vs, ds)
	}
}
