package str

import (
	"testing"
)

func TestIsAlNum(t *testing.T) {
	testData := map[string]bool{
		// data, want
		"0123456789":     true,
		"asodadnfas":     true,
		"fds0aenw0ew":    true,
		":a-daf9e0qfn":   false,
		":a-daf9e0\nqfn": false,
		"!#$%&'()0=~":    false,
		"Ω≈ç∫∂¨¨å…∂∑ei´˜ØÅÎÅÂ∏": false,
		"天上天下唯我独尊":              false,
		"mendakoめんだこ":           false,
		"":                      false,
	}
	for data, want := range testData {
		got := IsAlNum(data)
		if got != want {
			t.Errorf("data: %#v, want %#v, got: %#v", data, want, got)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	testData := map[string]bool{
		// data, want
		"0123456789":     false,
		"asodadnfas":     true,
		"fds0aenw0ew":    false,
		":a-daf9e0qfn":   false,
		":a-daf9e0\nqfn": false,
		"!#$%&'()0=~":    false,
		"Ω≈ç∫∂¨¨å…∂∑ei´˜ØÅÎÅÂ∏": false,
		"天上天下唯我独尊":              false,
		"mendakoめんだこ":           false,
		"":                      false,
	}
	for data, want := range testData {
		got := IsAlpha(data)
		if got != want {
			t.Errorf("data: %#v, want %#v, got: %#v", data, want, got)
		}
	}
}

func TestIsDigit(t *testing.T) {
	testData := map[string]bool{
		// data, want
		"0123456789":     true,
		"asodadnfas":     false,
		"fds0aenw0ew":    false,
		":a-daf9e0qfn":   false,
		":a-daf9e0\nqfn": false,
		"!#$%&'()0=~":    false,
		"Ω≈ç∫∂¨¨å…∂∑ei´˜ØÅÎÅÂ∏": false,
		"天上天下唯我独尊":              false,
		"mendakoめんだこ":           false,
		"":                      false,
	}
	for data, want := range testData {
		got := IsDigit(data)
		if got != want {
			t.Errorf("data: %#v, want %#v, got: %#v", data, want, got)
		}
	}
}

func TestIsASCII(t *testing.T) {
	testData := map[string]bool{
		// data, want
		"0123456789":     true,
		"asodadnfas":     true,
		"fds0aenw0ew":    true,
		":a-daf9e0qfn":   true,
		":a-daf9e0\nqfn": true,
		"!#$%&'()0=~":    true,
		"Ω≈ç∫∂¨¨å…∂∑ei´˜ØÅÎÅÂ∏": false,
		"天上天下唯我独尊":              false,
		"mendakoめんだこ":           false,
		"":                      true,
	}
	for data, want := range testData {
		got := IsASCII(data)
		if got != want {
			t.Errorf("data: %#v, want %#v, got: %#v", data, want, got)
		}
	}
}

func TestIsUpper(t *testing.T) {
	testData := map[string]bool{
		// data, want
		"0123456789":     false,
		"asodadnfas":     false,
		"fds0aenw0ew":    false,
		":a-daf9e0qfn":   false,
		":a-daf9e0\nqfn": false,
		"!#$%&'()0=~":    false,
		"Ω≈ç∫∂¨¨å…∂∑ei´˜ØÅÎÅÂ∏": false,
		"天上天下唯我独尊":              false,
		"mendakoめんだこ":           false,
		"":                      false,
		"sdfAIDF":               false,
		"AFSDMOI":               true,
		"FNSO-DFS.DFS":          true,
		"FDもずくFSO":              true,
		"sdffもずく":               false,
		"fasdf-sdfas.e":         false,
	}
	for data, want := range testData {
		got := IsUpper(data)
		if got != want {
			t.Errorf("data: %#v, want %#v, got: %#v", data, want, got)
		}
	}
}

func TestIsLower(t *testing.T) {
	testData := map[string]bool{
		// data, want
		"0123456789":     false,
		"asodadnfas":     true,
		"fds0aenw0ew":    true,
		":a-daf9e0qfn":   true,
		":a-daf9e0\nqfn": true,
		"!#$%&'()0=~":    false,
		"Ω≈ç∫∂¨¨å…∂∑ei´˜ØÅÎÅÂ∏": true,
		"天上天下唯我独尊":              false,
		"mendakoめんだこ":           true,
		"":                      false,
		"sdfAIDF":               false,
		"AFSDMOI":               false,
		"FNSO-DFS.DFS":          false,
		"FDもずくFSO":              false,
		"sdffもずく":               true,
		"fasdf-sdfas.e":         true,
	}
	for data, want := range testData {
		got := IsLower(data)
		if got != want {
			t.Errorf("data: %#v, want %#v, got: %#v", data, want, got)
		}
	}
}

func TestLen(t *testing.T) {
	testData := map[string]int{
		// data, want
		"もじゅく": 4,
		"∂≠å∂ƒ´˜ƒåπ∂ƒå´–ƒå∂": 18,
		"もっちもっちにゃんにゃん!mochi": 18,
		"天上天下唯我独尊":           8,
	}
	for data, want := range testData {
		got := Len(data)
		if got != want {
			t.Errorf("data: %#v, want %#v, got: %#v", data, want, got)
		}
	}
}
