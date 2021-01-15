// SyncStruct
// This module is similar to SyncFile.
// SyncStruct can synchronize a golang's struct data with a json file.
// Important:
// All fields in the struct, must start with upper case.

package syncFile

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

// This controller can stop the sync-goroutines
type SyncStructController struct {
	// The path to the file.
	filename string
	// Target struct data.
	structData interface{}
	// The interval to check for changes.
	interval time.Duration
	// If this flag is true, The sync-goroutines will be stopped.
	exitFlag bool
	// Modification time.
	modTime time.Time
	// Goroutine waiting group
	wg sync.WaitGroup
	// Latest error.
	err error
	// listener
	listener map[string]func(name string, old gjson.Result, new gjson.Result) error
	oldData map[string]gjson.Result
}

// Get filename
func (s *SyncStructController) Filename() string {
	return s.filename
}

// Get interval
func (s *SyncStructController) Interval() time.Duration {
	return s.interval
}

// Get modification time.
func (s *SyncStructController) ModTime() time.Time {
	return s.modTime
}

// Get latest error.
func (s *SyncStructController) Error() error {
	return s.err
}

// This goroutine is used to synchronize data.
func (s *SyncStructController) syncGoroutine() {
	timer := time.NewTicker(s.interval)
	defer timer.Stop()
	defer s.wg.Done()
	for range timer.C {
		// Check modification.
		// if info.ModTime() == s.modTime do nothing.
		if info, err := os.Stat(s.filename); err != nil {
			s.err = err
		} else if info.ModTime() != s.modTime {
			// detected modification, reload the file.
			_ = s.Reload()
		}
		// Exit flag.
		if s.exitFlag {
			return
		}
	}
}

// Reload the file manually.
func (s *SyncStructController) Reload() error {
	if info, err := os.Stat(s.filename); err != nil {
		s.err = err
		return err
	} else if contents, err := ioutil.ReadFile(s.filename); err != nil {
		s.err = err
		return err
	} else if contents != nil {
		// reload data
		if err := json.Unmarshal(contents, s.structData); err != nil {
			return err
		}
		s.modTime = info.ModTime()
		// check changes for listener
		res := gjson.ParseBytes(contents)
		for name := range s.listener {
			if value, ok := s.oldData[name]; ok && value.String() != res.Get(name).String(){
				err := s.listener[name](name, value, res.Get(name))
				if err != nil {
					s.err = err
				}
			}
			s.oldData[name] = res.Get(name)
		}
	}
	return nil
}

// Apply changes to file.
func (s *SyncStructController) Apply() error {
	if data, err := json.MarshalIndent(s.structData, "", "\t"); err != nil {
		return err
	} else if err := ioutil.WriteFile(s.filename, data, 0644); err != nil {
		return err
	} else {
		return nil
	}
}

// Stop sync-goroutines.
func (s *SyncStructController) Exit(wait bool) {
	s.exitFlag = true
	if wait {
		s.wg.Wait()
	}
}

// Add a event listener.
// When the value of this field in json file was changed, the listener function will be called.
// func(name string, old gjson.Result, new gjson.Result) error
// Important:
// name parameter is the field name in json file not in the struct.
// For example:
// type Sample struct {
// 	 Name string `json:"name"`   <--- This json tag.
// 	 Age int `json:"age"`
// }
// If you want to add a event listener to `Name` field, the name parameter must be "name"
func (s *SyncStructController) AddEventListener(name string, f func(string, gjson.Result, gjson.Result) error){
	s.listener[name] = f
}

// Create a goroutine to synchronize the data.
// If you want to stop the sync-goroutines, Use SyncStructController.Exit()
func SyncStructWithJsonFile(structData interface{}, filename string, interval time.Duration) (*SyncStructController, error) {
	controller := &SyncStructController{
		filename:   filename,
		structData: structData,
		interval:   interval,
		exitFlag:   false,
		listener: make(map[string]func(string, gjson.Result, gjson.Result)error),
		oldData: make(map[string]gjson.Result),
	}
	// Create file if not exist.
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	_ = file.Close()
	if data, err := ioutil.ReadFile(filename); err != nil {
		return nil, err
	} else if len(data) == 0 {
		if jsonData, err := json.MarshalIndent(controller.structData, "", "\t"); err != nil {
			return nil, err
		} else {
			if err := ioutil.WriteFile(filename, jsonData, 0644); err != nil {
				return nil, err
			}
		}
	}
	// Start sync.
	controller.wg.Add(1)
	go controller.syncGoroutine()
	return controller, nil
}
