package elarr

import "testing"

func TestUtils(t *testing.T) {
	if res := Sum([]interface{}{1, 2, 3}); res != 6 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := SumInt([]int{1, 2, 3}); res != 6 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := Average([]interface{}{1, 2, 3, 4, 5}); res != 3 {
		t.Errorf("Invalid result: %v", res)
	}
	if res := Join([]interface{}{1, 2, 3}, ","); res != "1,2,3" {
		t.Errorf("Invalid result: %v", res)
	}
	if IsEmpty([]interface{}{1, 2, 3}) {
		t.Errorf("Invalid result")
	}
	if !IsEmpty([]interface{}{}) {
		t.Errorf("Invalid result")
	}
}
