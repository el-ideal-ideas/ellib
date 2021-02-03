package elarr

import "testing"

func TestIndex(t *testing.T) {
	d := []interface{}{1, 2, 3, 4, 5, 3, 4, 5, 7}
	if res := IndexInter(d, 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexStr(ToStrings(d), "3"); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexInt(ToInts(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexInt8(ToInt8s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexInt16(ToInt16s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexInt32(ToInt32s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexInt64(ToInt64s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexUint(ToUints(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexUint8(ToUint8s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexUint16(ToUint16s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexUint32(ToUint32s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexUint64(ToUint64s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexRune(ToRunes(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexByte(ToBytes(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexFloat32(ToFloat32s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexFloat64(ToFloat64s(d), 3); res != 2 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := IndexBool(ToBools(d), false); res != -1 {
		t.Errorf("Invalid result: %v", res)
	}
}
