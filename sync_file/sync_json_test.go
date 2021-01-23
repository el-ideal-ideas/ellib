package sync_file

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestSyncJSON(t *testing.T) {
	filename := "./test.json"
	_ = os.Remove(filename)
	syncJson, err := NewSyncJSON(filename, time.Millisecond)
	if err != nil {
		t.Fatalf("Can't create instance of SyncJSON, Error: %s", err)
	}
	syncJson.Change([]byte(`
{
	"key00": "value",
	"key01": "もっちもっちにゃんにゃん",
	"key02": "いあいあめんだこちゃんふたぐん",
	"key03": [1, 2, 3, 4],
	"key04": ["A", "B", "C", "D"],
	"key05": 1024,
	"key06": 3.1415,
	"key07": false,
	"key08": 1,
	"key09": true
}
	`))
	time.Sleep(time.Second)
	if d := syncJson.Get("key00").String(); d != "value" {
		t.Errorf("Invalid data: %#v", d)
	}
	if d := syncJson.Get("key01").String(); d != "もっちもっちにゃんにゃん" {
		t.Errorf("Invalid data: %#v", d)
	}
	if d := syncJson.Get("key02").String(); d != "いあいあめんだこちゃんふたぐん" {
		t.Errorf("Invalid data: %#v", d)
	}
	if d := syncJson.Get("key03").String(); d != "[1, 2, 3, 4]" {
		t.Errorf("Invalid data: %#v", d)
	}
	if d := syncJson.Get("key04").String(); d != "[\"A\", \"B\", \"C\", \"D\"]" {
		t.Errorf("Invalid data: %#v", d)
	}
	if d := syncJson.Get("key05").String(); d != "1024" {
		t.Errorf("Invalid data: %#v", d)
	}
	if d := syncJson.GetInt("key05"); d != 1024 {
		t.Errorf("The response must be 1024, but got %#v", d)
	}
	if d := syncJson.GetInt("key00"); d != 0 {
		t.Errorf("The response must be 0, but got %#v", d)
	}
	if d := syncJson.GetInt("key01"); d != 0 {
		t.Errorf("The response must be 0, but got %#v", d)
	}
	if d := syncJson.GetInt("key03"); d != 0 {
		t.Errorf("The response must be 0, but got %#v", d)
	}
	if d := syncJson.GetInt("key06"); d != 3 {
		t.Errorf("The response must be 3, but got %#v", d)
	}
	if d := syncJson.GetString("key00"); d != "value" {
		t.Errorf("The response must be \"value\", but got %#v", d)
	}
	if d := syncJson.GetString("key03"); d != "[1, 2, 3, 4]" {
		t.Errorf("The response must be \"[1, 2, 3, 4]\", but got %#v", d)
	}
	if d := syncJson.GetFloat("key03"); d != 0.0 {
		t.Errorf("The response must be 0.0, but got %#v", d)
	}
	if d := syncJson.GetFloat("key05"); d != 1024.0 {
		t.Errorf("The response must be 1024.0, but got %#v", d)
	}
	if d := syncJson.GetFloat("key06"); d != 3.1415 {
		t.Errorf("The response must be 3.1415, but got %#v", d)
	}
	if d := syncJson.GetBool("key00"); d != false {
		t.Errorf("The response must be false, but got %#v", d)
	}
	if d := syncJson.GetBool("key07"); d != false {
		t.Errorf("The response must be false, but got %#v", d)
	}
	if d := syncJson.GetBool("key08"); d != true {
		t.Errorf("The response must be true, but got %#v", d)
	}
	if d := syncJson.GetBool("key09"); d != true {
		t.Errorf("The response must be true, but got %#v", d)
	}
	if d := syncJson.GetArray("key03"); !(d[0].Int() == 1 && d[1].Int() == 2 && d[2].Int() == 3 && d[3].Int() == 4) {
		t.Errorf("The response must be [1, 2, 3, 4], but got %#v", d)
	}
	if d := syncJson.GetArrayOfInt("key03"); !(d[0] == 1 && d[1] == 2 && d[2] == 3 && d[3] == 4) {
		t.Errorf("The response must be [1, 2, 3, 4], but got %#v", d)
	}
	if d := syncJson.GetArrayOfString("key04"); !(d[0] == "A" && d[1] == "B" && d[2] == "C" && d[3] == "D") {
		t.Errorf("The response must be [\"A\", \"B\", \"C\", \"D\"], but got %#v", d)
	}
	// --------
	if err := ioutil.WriteFile(filename, []byte(`
{
	"key00": "value",
	"key01": "もっちもっちにゃんにゃん!!!!!!!!!!",
	"key02": "いあいあめんだこちゃんふたぐん",
	"key03": [1, 2, 3, 4],
	"key04": ["A", "B", "C", "D"],
	"key05": 1024,
	"key06": 3.1415,
	"key07": false,
	"key08": 1,
	"key09": true
}
	`), 0644); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	time.Sleep(time.Second)
	if d := syncJson.GetString("key01"); d != "もっちもっちにゃんにゃん!!!!!!!!!!" {
		t.Errorf("Invalid data: %#v", d)
	}
	// --------
	if err := syncJson.Reload(); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	syncJson.RunCommand(cmdReload)
	syncJson.RunCommand(cmdWrite)
	if syncJson.Error() != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	syncJson.Exit(true)
	if syncJson.Error() != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	_ = os.Remove(filename)
}
