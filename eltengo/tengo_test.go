package eltengo

import (
	"testing"
)

func TestTengo(t *testing.T) {
	executor := New("result := a + b")
	result, err := executor(map[string]interface{}{
		"a": "もずく",
		"b": "美味しい",
	})
	if err != nil {
		t.Error(err)
	} else if result.String() != "もずく美味しい" {
		t.Errorf("Invalid result.")
	}
	result, err = executor(map[string]interface{}{
		"a": 20,
		"b": 7,
	})
	if err != nil {
		t.Error(err)
	} else if result.Int() != 27 {
		t.Errorf("Invalid result.")
	}
}
