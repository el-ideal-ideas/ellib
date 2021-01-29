package elsyncfile

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestSyncFile(t *testing.T) {
	filename := "./test.txt"
	_ = os.Remove(filename)
	syncFile, err := NewSyncFile(filename, time.Millisecond)
	// --------
	if err != nil {
		t.Fatalf("Can't create instance of SyncFile, Error: %s", err)
	}
	if f := syncFile.Filename(); f != syncFile.filename {
		t.Errorf("Getter returned an invalid value, Filename: %s", f)
	}
	if i := syncFile.Interval(); i != syncFile.interval {
		t.Errorf("Getter returned an invalid value, Interval: %s", i)
	}
	if c := syncFile.Contents(); !bytes.Equal(c, syncFile.contents) {
		t.Errorf("Getter returned an invalid value, Contents: %s", c)
	}
	if c := syncFile.Control(); c != syncFile.control {
		t.Errorf("Getter returned an invalid value, Control: %v", c)
	}
	if m := syncFile.ModTime(); m != syncFile.modTime {
		t.Errorf("Getter returned an invalid value, ModTime: %s", m)
	}
	if e := syncFile.Error(); e != syncFile.err {
		t.Errorf("Getter returned an invalid value, Error: %s", e)
	}
	// --------
	if !bytes.Equal(syncFile.Contents(), make([]byte, 0)) {
		t.Errorf("Contents() must be an empty slice of byte.")
	}
	if syncFile.GetString() != "" {
		t.Errorf("GetString() must be an empty string.")
	}
	if info, err := os.Stat(filename); err != nil {
		t.Errorf("Can't get the file info. Error: %s", err)
	} else if info.ModTime() != syncFile.ModTime() {
		t.Errorf("Invalid value of ModTime()")
	}
	// --------
	if err := syncFile.Change([]byte("もっちもっちにゃんにゃん")); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if !bytes.Equal(syncFile.Contents(), []byte("もっちもっちにゃんにゃん")) {
		t.Errorf("Contents() must be []byte(\"もっちもっちにゃんにゃん\").")
	}
	if syncFile.GetString() != "もっちもっちにゃんにゃん" {
		t.Errorf("GetString() must be \"もっちもっちにゃんにゃん\".")
	}
	if data, err := ioutil.ReadFile(filename); err != nil {
		t.Errorf("Expected nil, got %v", err)
	} else if !bytes.Equal(data, syncFile.Contents()) {
		t.Errorf("Invalid contents")
	}
	// --------
	if err := ioutil.WriteFile(filename, []byte("いあいあめんだこちゃんふたぐん"), os.ModePerm); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	time.Sleep(time.Second)
	if !bytes.Equal(syncFile.Contents(), []byte("いあいあめんだこちゃんふたぐん")) {
		t.Errorf("Contents() must be []byte(\"いあいあめんだこちゃんふたぐん\").")
	}
	if syncFile.GetString() != "いあいあめんだこちゃんふたぐん" {
		t.Errorf("GetString() must be \"いあいあめんだこちゃんふたぐん\".")
	}
	if info, err := os.Stat(filename); err != nil {
		t.Errorf("Can't get the file info. Error: %s", err)
	} else if info.ModTime() != syncFile.ModTime() {
		t.Errorf("Invalid value of ModTime()")
	}
	// --------
	if err := syncFile.Reload(); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	syncFile.RunCommand(cmdReload)
	syncFile.RunCommand(cmdWrite)
	if syncFile.Error() != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	syncFile.Exit(true)
	if syncFile.Error() != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	_ = os.Remove(filename)
}
