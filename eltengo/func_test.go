package eltengo

import "testing"

func TestCalc(t *testing.T) {
	if res, err := Calc("1+2"); err != nil {
		t.Error(err)
	} else if res != 3 {
		t.Errorf("Invalid result: %f", res)
	}
	if res, err := Calc("1+2*3"); err != nil {
		t.Error(err)
	} else if res != 7 {
		t.Errorf("Invalid result: %f", res)
	}
	if res, err := Calc("1+2*3.2"); err != nil {
		t.Error(err)
	} else if res != 7.4 {
		t.Errorf("Invalid result: %f", res)
	}
}
