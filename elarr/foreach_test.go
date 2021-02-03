package elarr

import "testing"

func TestForeach(t *testing.T) {
	d := []interface{}{1, 2, 3, 4, 5}
	i := 1
	ForEachInter(d, func(v interface{}) { i *= v.(int) })
	if i != 120 {
		t.Errorf("Invalid result")
	}
	j := ""
	ForEachStr(ToStrings(d), func(v string) { j += v })
	if j != "12345" {
		t.Errorf("Invalid result")
	}
	var iv int
	ForEachInt(ToInts(d), func(v int) { iv += v })
	if iv != 15 {
		t.Errorf("Invalid result")
	}
	var iv8 int8
	ForEachInt8(ToInt8s(d), func(v int8) { iv8 += v })
	if iv8 != 15 {
		t.Errorf("Invalid result")
	}
	var iv16 int16
	ForEachInt16(ToInt16s(d), func(v int16) { iv16 += v })
	if iv16 != 15 {
		t.Errorf("Invalid result")
	}
	var iv32 int32
	ForEachInt32(ToInt32s(d), func(v int32) { iv32 += v })
	if iv32 != 15 {
		t.Errorf("Invalid result")
	}
	var iv64 int64
	ForEachInt64(ToInt64s(d), func(v int64) { iv64 += v })
	if iv64 != 15 {
		t.Errorf("Invalid result")
	}
	var uiv uint
	ForEachUint(ToUints(d), func(v uint) { uiv += v })
	if uiv != 15 {
		t.Errorf("Invalid result")
	}
	var uiv8 uint8
	ForEachUint8(ToUint8s(d), func(v uint8) { uiv8 += v })
	if uiv8 != 15 {
		t.Errorf("Invalid result")
	}
	var uiv16 uint16
	ForEachUint16(ToUint16s(d), func(v uint16) { uiv16 += v })
	if uiv16 != 15 {
		t.Errorf("Invalid result")
	}
	var uiv32 uint32
	ForEachUint32(ToUint32s(d), func(v uint32) { uiv32 += v })
	if uiv32 != 15 {
		t.Errorf("Invalid result")
	}
	var uiv64 uint64
	ForEachUint64(ToUint64s(d), func(v uint64) { uiv64 += v })
	if uiv64 != 15 {
		t.Errorf("Invalid result")
	}
	var r rune
	ForEachRune(ToRunes(d), func(v rune) { r += v })
	if r != 15 {
		t.Errorf("Invalid result")
	}
	var b byte
	ForEachByte(ToBytes(d), func(v byte) { b += v })
	if b != 15 {
		t.Errorf("Invalid result")
	}
	var f32 float32
	ForEachFloat32(ToFloat32s(d), func(v float32) { f32 += v })
	if f32 != 15 {
		t.Errorf("Invalid result")
	}
	var f64 float64
	ForEachFloat64(ToFloat64s(d), func(v float64) { f64 += v })
	if f64 != 15 {
		t.Errorf("Invalid result")
	}
	var bo int
	ForEachBool(ToBools(d), func(v bool) {
		if v {
			bo += 1
		}
	})
	if bo != 5 {
		t.Errorf("Invalid result")
	}
}
