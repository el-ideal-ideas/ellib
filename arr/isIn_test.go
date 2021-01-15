package arr

import "testing"


func TestIsIn(t *testing.T) {
	if !StrIn([]string{"も", "ず", "く"}, "も") {
		t.Errorf("Invalid result")
	}
	if StrIn([]string{"も", "ず", "く"}, "あ") {
		t.Errorf("Invalid result")
	}
	if !IntIn([]int{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if IntIn([]int{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Int8In([]int8{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if Int8In([]int8{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Int16In([]int16{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if Int16In([]int16{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Int32In([]int32{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if Int32In([]int32{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Int64In([]int64{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if Int64In([]int64{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !UintIn([]uint{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if UintIn([]uint{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Uint8In([]uint8{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if Uint8In([]uint8{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Uint16In([]uint16{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if Uint16In([]uint16{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Uint32In([]uint32{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if Uint32In([]uint32{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Uint64In([]uint64{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if Uint64In([]uint64{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !RuneIn([]rune{'も', 'ず', 'く'}, 'も') {
		t.Errorf("Invalid result")
	}
	if RuneIn([]rune{'も', 'ず', 'く'}, 'あ') {
		t.Errorf("Invalid result")
	}
	if !ByteIn([]byte{1, 2, 3}, 1) {
		t.Errorf("Invalid result")
	}
	if ByteIn([]byte{1, 2, 3}, 0) {
		t.Errorf("Invalid result")
	}
	if !Float32In([]float32{1.0, 2.0, 3.0}, 1.0) {
		t.Errorf("Invalid result")
	}
	if Float32In([]float32{1.0, 2.0, 3.0}, 0.0) {
		t.Errorf("Invalid result")
	}

}
