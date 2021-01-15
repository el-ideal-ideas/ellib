// SyncJSON
// This module is similar to SyncFile module, But will be treated the file as a json file.
// Json data in the memory will be synchronized with the file.

package syncFile

import (
	"github.com/tidwall/gjson"
	"os"
	"time"
)


// A json dictionary file object that will be synchronized with the target file.
// If you change the file directly, SyncJSON will reload contents from the file.
// If you change the contents in SyncJSON, changes will be applied to the file.
// Warning:
// All contents in the file will be loaded into memory, Please be careful to the size of the file.
type SyncJSON struct {
	SyncFile
	Json gjson.Result
}

// Get value from dict.
func (s *SyncJSON) Get(key string) gjson.Result {
	return s.Json.Get(key)
}

// Get value as int.
func (s *SyncJSON) GetInt(key string) int {
	return int(s.Json.Get(key).Int())
}

// Get value as string.
func (s *SyncJSON) GetString(key string) string {
	return s.Json.Get(key).String()
}

// Get value as float.
func (s *SyncJSON) GetFloat(key string) float64 {
	return s.Json.Get(key).Float()
}

// Get value as bool.
func (s *SyncJSON) GetBool(key string) bool {
	return s.Json.Get(key).Bool()
}

// Get value as array.
func (s *SyncJSON) GetArray(key string) []gjson.Result {
	return s.Json.Get(key).Array()
}

// Get value as a array of string.
func (s *SyncJSON) GetArrayOfString(key string) []string {
	arr := s.Json.Get(key).Array()
	ret := make([]string, len(arr))
	for index, item := range arr {
		ret[index] = item.String()
	}
	return ret
}

// Get value as a array of int.
func (s *SyncJSON) GetArrayOfInt(key string) []int {
	arr := s.Json.Get(key).Array()
	ret := make([]int, len(arr))
	for index, item := range arr {
		ret[index] = int(item.Int())
	}
	return ret
}

// Try return JSON as array
func (s *SyncJSON) Array() []gjson.Result {
	return s.Json.Array()
}

// Try return JSON as dict
func (s *SyncJSON) Map() map[string]gjson.Result {
	return s.Json.Map()
}

// Create a SyncJSON object.
func NewSyncJSON(filename string, interval time.Duration) (*SyncJSON, error) {
	syncJSON := &SyncJSON{
		SyncFile: SyncFile{
			filename: filename,
			interval: interval,
			control:  make(chan command, 32),
		},
	}
	// Open file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Load File
	if _, err = file.Read(syncJSON.contents); err != nil {
		return nil, err
	}
	// Get Info
	info, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	syncJSON.modTime = info.ModTime()
	// Reload function
	syncJSON.reloadFunc = func(i *SyncFile){
		syncJSON.Json = gjson.ParseBytes(syncJSON.contents)
	}
	// Start loop-goroutine
	syncJSON.wg.Add(1)
	go syncJSON.loopGoroutine()
	// Parse JSON
	syncJSON.Json = gjson.ParseBytes(syncJSON.contents)
	return syncJSON, nil
}