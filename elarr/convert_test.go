package elarr

import "testing"

func TestConvert(t *testing.T) {
	if data := ToStrings([]interface{}{"も", "く", "ず"}); !IsSameStr(data, []string{"も", "く", "ず"}) {
		t.Errorf("Invalid result")
	}
	if data := ToInts([]interface{}{1, 2, 3}); !IsSameInt(data, []int{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToInt8s([]interface{}{int8(1), int8(2), int8(3)}); !IsSameInt8(data, []int8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToInt16s([]interface{}{int16(1), int16(2), int16(3)}); !IsSameInt16(data, []int16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToInt32s([]interface{}{int32(1), int32(2), int32(3)}); !IsSameInt32(data, []int32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToInt64s([]interface{}{int64(1), int64(2), int64(3)}); !IsSameInt64(data, []int64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToUints([]interface{}{uint(1), uint(2), uint(3)}); !IsSameUint(data, []uint{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToUint8s([]interface{}{uint8(1), uint8(2), uint8(3)}); !IsSameUint8(data, []uint8{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToUint16s([]interface{}{uint16(1), uint16(2), uint16(3)}); !IsSameUint16(data, []uint16{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToUint32s([]interface{}{uint32(1), uint32(2), uint32(3)}); !IsSameUint32(data, []uint32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToUint64s([]interface{}{uint64(1), uint64(2), uint64(3)}); !IsSameUint64(data, []uint64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToBytes([]interface{}{byte(1), byte(2), byte(3)}); !IsSameByte(data, []byte{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToRunes([]interface{}{rune(1), rune(2), rune(3)}); !IsSameRune(data, []rune{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToFloat32s([]interface{}{float32(1), float32(2), float32(3)}); !IsSameFloat32(data, []float32{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := ToFloat64s([]interface{}{float64(1), float64(2), float64(3)}); !IsSameFloat64(data, []float64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := IntsToStrings([]int{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Int8sToStrings([]int8{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Int16sToStrings([]int16{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Int32sToStrings([]int32{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Int64sToStrings([]int64{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := UintsToStrings([]uint{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Uint8sToStrings([]uint8{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Uint16sToStrings([]uint16{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Uint32sToStrings([]uint32{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Uint64sToStrings([]uint64{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := RunesToStrings([]rune{'も', 'ず', 'く'}); !IsSameStr(data, []string{"も", "ず", "く"}) {
		t.Errorf("Invalid result")
	}
	if data := BytesToStrings([]byte{1, 2, 3}); !IsSameStr(data, []string{"1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Float32sToStrings([]float32{1.1, 2, 3}, 'f', -1, 32); !IsSameStr(data, []string{"1.1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := Float64sToStrings([]float64{1.1, 2, 3}, 'f', -1, 32); !IsSameStr(data, []string{"1.1", "2", "3"}) {
		t.Errorf("Invalid result")
	}
	if data := IntsToFloat64s([]int{1, 2, 3}); !IsSameFloat64(data, []float64{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
	if data := Float64sToInts([]float64{1.0, 2.0, 3.0}); !IsSameInt(data, []int{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := IntsToUints([]int{1, 2, 3, -10}); !IsSameUint(data, []uint{1, 2, 3, 0}) {
		t.Errorf("Invalid result")
	}
	if data := UintsToInts([]uint{1, 2, 3}); !IsSameInt(data, []int{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := UintsToFloat64s([]uint{1, 2, 3}); !IsSameFloat64(data, []float64{1.0, 2.0, 3.0}) {
		t.Errorf("Invalid result")
	}
	if data := Float64sToUints([]float64{1.0, 2.0, 3.0}); !IsSameUint(data, []uint{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data, _ := StringsToInts([]string{"1", "2", "3"}); !IsSameInt(data, []int{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data, _ := StringsToUints([]string{"1", "2", "3", "-1"}); !IsSameUint(data, []uint{1, 2, 3, 0}) {
		t.Errorf("Invalid result")
	}
	if data, _ := StringsToFloat64s([]string{"1", "2", "3"}); !IsSameFloat64(data, []float64{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if data := RunesToBytes([]rune{'も', 'ず', 'く'}); !IsSameByte(data, []byte("もずく")) {
		t.Errorf("Invalid result")
	}
}
