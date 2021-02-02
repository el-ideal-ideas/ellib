package elarr

import "testing"

func TestUnique(t *testing.T) {
	data := []interface{}{1, 2, 3, 2}
	want := []interface{}{1, 2, 3}
	if res := UniqueInter(data); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueStr(ToStrings(data)); !IsSame(res, ToStrings(want)) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueInt(ToInts(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueInt8(ToInt8s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueInt16(ToInt16s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueInt32(ToInt32s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueInt64(ToInt64s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueUint(ToUints(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueUint8(ToUint8s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueUint16(ToUint16s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueUint32(ToUint32s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueUint64(ToUint64s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueFloat32(ToFloat32s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueFloat64(ToFloat64s(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueRune(ToRunes(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
	if res := UniqueByte(ToBytes(data)); !IsSame(res, want) {
		t.Errorf("Invalid result: %v", res)
	}
}
