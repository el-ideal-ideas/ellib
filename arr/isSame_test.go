package arr

import "testing"

func TestIsSame(t *testing.T) {
	if !StrIsSame([]string{"も", "ず", "く"}, []string{"も", "ず", "く"}, []string{"も", "ず", "く"}) {
		t.Errorf("Invalid result")
	}
	if StrIsSame([]string{"あ", "ず", "く"}, []string{"も", "ず", "く"}, []string{"も", "ず", "く"}) {
		t.Errorf("Invalid result")
	}
	if !IntIsSame([]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if IntIsSame([]int{0, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Int8IsSame([]int8{1, 2, 3}, []int8{1, 2, 3}, []int8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if Int8IsSame([]int8{0, 2, 3}, []int8{1, 2, 3}, []int8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Int16IsSame([]int16{1, 2, 3}, []int16{1, 2, 3}, []int16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if Int16IsSame([]int16{0, 2, 3}, []int16{1, 2, 3}, []int16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Int32IsSame([]int32{1, 2, 3}, []int32{1, 2, 3}, []int32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if Int32IsSame([]int32{0, 2, 3}, []int32{1, 2, 3}, []int32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Int64IsSame([]int64{1, 2, 3}, []int64{1, 2, 3}, []int64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if Int64IsSame([]int64{0, 2, 3}, []int64{1, 2, 3}, []int64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !UintIsSame([]uint{1, 2, 3}, []uint{1, 2, 3}, []uint{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if UintIsSame([]uint{0, 2, 3}, []uint{1, 2, 3}, []uint{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Uint8IsSame([]uint8{1, 2, 3}, []uint8{1, 2, 3}, []uint8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if Uint8IsSame([]uint8{0, 2, 3}, []uint8{1, 2, 3}, []uint8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Uint16IsSame([]uint16{1, 2, 3}, []uint16{1, 2, 3}, []uint16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if Uint16IsSame([]uint16{0, 2, 3}, []uint16{1, 2, 3}, []uint16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Uint32IsSame([]uint32{1, 2, 3}, []uint32{1, 2, 3}, []uint32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if Uint32IsSame([]uint32{0, 2, 3}, []uint32{1, 2, 3}, []uint32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Uint64IsSame([]uint64{1, 2, 3}, []uint64{1, 2, 3}, []uint64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if Uint64IsSame([]uint64{0, 2, 3}, []uint64{1, 2, 3}, []uint64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !RuneIsSame([]rune{'も', 'ず', 'く'}, []rune{'も', 'ず', 'く'}, []rune{'も', 'ず', 'く'}) {
		t.Errorf("Invalid result")
	}
	if RuneIsSame([]rune{'あ', 'ず', 'く'}, []rune{'も', 'ず', 'く'}, []rune{'も', 'ず', 'く'}) {
		t.Errorf("Invalid result")
	}
	if !ByteIsSame([]byte{1, 2, 3}, []byte{1, 2, 3}, []byte{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if ByteIsSame([]byte{0, 2, 3}, []byte{1, 2, 3}, []byte{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !Float32IsSame([]float32{1.0, 2.0, 3.0}, []float32{1.0, 2.0, 3.0}, []float32{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
	if Float32IsSame([]float32{0.0, 2.0, 3.0}, []float32{1.0, 2.0, 3.0}, []float32{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
	if !Float64IsSame([]float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
	if Float64IsSame([]float64{0.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
}
