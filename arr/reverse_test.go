package arr

import "testing"


func TestReverse(t *testing.T) {
	d1 := []string{"も", "ず", "く"}
	StrReverse(d1)
	if !StrIsSame(d1, []string{"く", "ず", "も"}) {
		t.Errorf("Invalid result")
	}
	d2 := []int{1, 2, 3}
	IntReverse(d2)
	if !IntIsSame(d2, []int{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d3 := []int8{1, 2, 3}
	Int8Reverse(d3)
	if !Int8IsSame(d3, []int8{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d4 := []int16{1, 2, 3}
	Int16Reverse(d4)
	if !Int16IsSame(d4, []int16{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d5 := []int32{1, 2, 3}
	Int32Reverse(d5)
	if !Int32IsSame(d5, []int32{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d6 := []int64{1, 2, 3}
	Int64Reverse(d6)
	if !Int64IsSame(d6, []int64{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d7 := []uint{1, 2, 3}
	UintReverse(d7)
	if !UintIsSame(d7, []uint{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d8 := []uint8{1, 2, 3}
	Uint8Reverse(d8)
	if !Uint8IsSame(d8, []uint8{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d9 := []uint16{1, 2, 3}
	Uint16Reverse(d9)
	if !Uint16IsSame(d9, []uint16{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d10 := []uint32{1, 2, 3}
	Uint32Reverse(d10)
	if !Uint32IsSame(d10, []uint32{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d11 := []uint64{1, 2, 3}
	Uint64Reverse(d11)
	if !Uint64IsSame(d11, []uint64{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d12 := []rune{'も', 'ず', 'く'}
	RuneReverse(d12)
	if !RuneIsSame(d12, []rune{'く', 'ず', 'も'}) {
		t.Errorf("Invalid result")
	}
	d13 := []byte{1, 2, 3}
	ByteReverse(d13)
	if !ByteIsSame(d13, []byte{3, 2, 1}) {
		t.Errorf("Invalid result")
	}
	d14 := []float32{1.0, 2.0, 3.0}
	Float32Reverse(d14)
	if !Float32IsSame(d14, []float32{3.0, 2.0, 1.0}) {
		t.Errorf("Invalid result")
	}
	d15 := []float64{1.0, 2.0, 3.0}
	Float64Reverse(d15)
	if !Float64IsSame(d15, []float64{3.0, 2.0, 1.0}) {
		t.Errorf("Invalid result")
	}
}
