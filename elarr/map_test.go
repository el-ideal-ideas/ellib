package elarr

import (
	"github.com/el-ideal-ideas/ellib/elconv"
	"github.com/el-ideal-ideas/ellib/elstr"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	s := []interface{}{"も", "ず", "く"}
	sr := []interface{}{"もも", "ずず", "くく"}
	i := []interface{}{1, 2, 3}
	ir := []interface{}{2, 4, 6}
	ir2 := []interface{}{"1", "2", "3"}
	if d := MapStrInter(ToStrings(s), func(s string) interface{} { return s + s }); !IsSame(d, sr) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapStrStr(ToStrings(s), func(s string) string { return s + s }); !IsSame(d, sr) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapStrInt(ToStrings(s), func(s string) int { return elstr.Len(s) }); !IsSame(d, []int{1, 1, 1}) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapStrUint(ToStrings(s), func(s string) uint { return uint(elstr.Len(s)) }); !IsSame(d, []uint{1, 1, 1}) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapStrFloat64(ToStrings(s), func(s string) float64 { return float64(elstr.Len(s)) }); !IsSame(d, []float64{1, 1, 1}) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapIntInter(ToInts(i), func(v int) interface{} { return v * 2 }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapIntStr(ToInts(i), func(v int) string { return strconv.Itoa(v) }); !IsSame(d, ir2) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapIntInt(ToInts(i), func(v int) int { return v * 2 }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapIntUint(ToInts(i), func(v int) uint { return uint(v * 2) }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapIntFloat64(ToInts(i), func(v int) float64 { return float64(v * 2) }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapFloat64Inter(ToFloat64s(i), func(v float64) interface{} { return v * 2 }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapFloat64Str(ToFloat64s(i), func(v float64) string { return elconv.AsStr(v) }); !IsSame(d, ir2) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapFloat64Int(ToFloat64s(i), func(v float64) int { return int(v * 2) }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapFloat64Uint(ToFloat64s(i), func(v float64) uint { return uint(v * 2) }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapFloat64Float64(ToFloat64s(i), func(v float64) float64 { return v * 2 }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapUintInter(ToUints(i), func(v uint) interface{} { return v * 2 }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapUintStr(ToUints(i), func(v uint) string { return strconv.Itoa(int(v)) }); !IsSame(d, ir2) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapUintInt(ToUints(i), func(v uint) int { return int(v * 2) }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapUintUint(ToUints(i), func(v uint) uint { return v * 2 }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapUintFloat64(ToUints(i), func(v uint) float64 { return float64(v * 2) }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapInterInter(i, func(v interface{}) interface{} { return v.(int) * 2 }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapInterStr(i, func(v interface{}) string { return strconv.Itoa(v.(int)) }); !IsSame(d, ir2) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapInterInt(i, func(v interface{}) int { return v.(int) * 2 }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapInterUint(i, func(v interface{}) uint { return uint(v.(int) * 2) }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
	if d := MapInterFloat64(i, func(v interface{}) float64 { return float64(v.(int) * 2) }); !IsSame(d, ir) {
		t.Errorf("Invalid result: %v", d)
	}
}
