package elarr

import "testing"

func TestLIndex(t *testing.T) {
	d := []interface{}{1, 2, 3, 4, 5, 3, 4, 5, 7}
	if res := LIndexInter(d, 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexStr(ToStrings(d), "3"); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexInt(ToInts(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexInt8(ToInt8s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexInt16(ToInt16s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexInt32(ToInt32s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexInt64(ToInt64s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexUint(ToUints(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexUint8(ToUint8s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexUint16(ToUint16s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexUint32(ToUint32s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexUint64(ToUint64s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexRune(ToRunes(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexByte(ToBytes(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexFloat32(ToFloat32s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexFloat64(ToFloat64s(d), 3); res != 5 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := LIndexBool(ToBools(d), false); res != -1 {
		t.Errorf("Invalid result: %v", res)
	}
}
