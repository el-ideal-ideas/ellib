package elsyncfile

import (
	"os"
	"strings"
	"testing"
	"time"
)

type Tmp struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	List []int  `json:"list"`
}

func TestSyncStruct(t *testing.T) {
	filename := "./struct.json"
	_ = os.Remove(filename)
	json := `{
	"name": "もずく",
	"age": 18,
	"list": [
		1,
		2,
		3,
		4
	]
}`
	structData := Tmp{
		Name: "もずく",
		Age:  18,
		List: []int{1, 2, 3, 4},
	}
	controller, err := SyncStructWithJsonFile(&structData, filename, time.Millisecond)
	if err != nil {
		t.Fatalf("Can't create instance of SyncStructController, Error: %v", err)
	}
	if data, err := os.ReadFile(filename); err != nil {
		t.Errorf("Can't access to the file. Error: %v", err)
	} else if string(data) != json {
		t.Errorf("Invalid file contents.")
	}
	if structData.Name != "もずく" {
		t.Errorf("structData.Name must be \"もずく\"")
	}
	structData.Age = 19
	if err := controller.Apply(); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if data, err := os.ReadFile(filename); err != nil {
		t.Errorf("Can't access to the file. Error: %v", err)
	} else if string(data) != strings.ReplaceAll(json, "18", "19") {
		t.Errorf("Invalid file contents.")
	}
	if err := os.WriteFile(filename, []byte(strings.ReplaceAll(json, "もずく", "もじゅく")), 0644); err != nil {
		t.Errorf("Expected nil, got %v", err)
	} else {
		time.Sleep(time.Second)
		if structData.Name != "もじゅく" {
			t.Errorf("structData.Name must be \"もじゅく\"")
		}
	}
	if err := os.WriteFile(filename, []byte(strings.ReplaceAll(json, "4", "5")), 0644); err != nil {
		t.Errorf("Expected nil, got %v", err)
	} else {
		time.Sleep(time.Second)
		if structData.List[3] != 5 {
			t.Errorf("structData.List[3] must be 5")
		}
	}
	if err := controller.Reload(); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if controller.Error() != nil {
		t.Errorf("Expected nil, got %v", controller.Error())
	}
	controller.Exit(true)
	if controller.Error() != nil {
		t.Errorf("Expected nil, got %v", controller.Error())
	}
	_ = os.Remove(filename)
}
