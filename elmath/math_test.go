package elmath

import "testing"

func TestRound(t *testing.T) {
	if res := Round(1.2345, 1); res != 1.2 {
		t.Errorf("Invalid result: %f", res)
	}
	if res := Round(1.2345, 3); res != 1.235 {
		t.Errorf("Invalid result: %f", res)
	}
	if res := Percent(1, 2); res != 50 {
		t.Errorf("Invalid result: %f", res)
	}
	if res := Factorial(5); res != 120 {
		t.Errorf("Invalid result: %d", res)
	}
	if res := Permutation(5, 3); res != 60 {
		t.Errorf("Invalid result: %d", res)
	}
	if res := Combination(5, 3); res != 10 {
		t.Errorf("Invalid result: %d", res)
	}
}
