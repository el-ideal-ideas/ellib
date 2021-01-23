// SyncFile
// This module can synchronize data between memory and file.
// If you change the file, The new data will be reloaded to memory automatically.
// If you change the data in memory, Changes will be applied to file automatically.

package syncFile

import (
	"io/ioutil"
	"os"
	"sync"
	"time"
)

// A file object that will be synchronized with the target file.
// If you change the file directly, SyncFile will reload contents from the file.
// If you change the contents in SyncFile, changes will be applied to the file.
// Warning:
// All contents in the file will be loaded into memory, Please be careful to the size of the file.
type SyncFile struct {
	// The path to the file.
	filename string
	// The interval to check for changes.
	interval time.Duration
	// The contents of the file.
	contents []byte
	// control commands
	control chan command
	// Modification time.
	modTime time.Time
	// Latest error.
	err error
	// Goroutine waiting group
	wg sync.WaitGroup
	// This is a function to customize reload process.
	reloadFunc func(*SyncFile)
}

// Commands for loop-goroutine
type command uint32

// Commands for the loop-goroutine.
const (
	cmdReload command = iota
	cmdWrite
	cmdExit
)

// Get filename
func (s *SyncFile) Filename() string {
	return s.filename
}

// Get interval
func (s *SyncFile) Interval() time.Duration {
	return s.interval
}

// Access to contents directly
// Important:
// If you changed the contents directly, You need run the Apply() function to apply changes.
func (s *SyncFile) Contents() []byte {
	return s.contents
}

// Get contents as string.
func (s *SyncFile) GetString() string {
	return string(s.contents)
}

// Access to control channel directly.
func (s *SyncFile) Control() chan command {
	return s.control
}

// Get modification time.
func (s *SyncFile) ModTime() time.Time {
	return s.modTime
}

// Get latest error.
func (s *SyncFile) Error() error {
	return s.err
}

// The loop-goroutine
// This goroutine will execute commands in SyncFile.control
// And check the file modification time. If the file was modified, reload the contents.
func (s *SyncFile) loopGoroutine() {
	timer := time.NewTicker(s.interval)
	defer timer.Stop()
	defer close(s.control)
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
		// Check control commands.
		select {
		case cmd := <-s.control:
			if cmd == cmdExit {
				return
			} else {
				s.RunCommand(cmd)
			}
		default:
			// do nothing.
		}
	}
}

// Run commands.
func (s *SyncFile) RunCommand(cmd command) {
	switch cmd {
	case cmdReload:
		_ = s.Reload()
	case cmdWrite:
		_ = s.Apply()
	}
}

// Reload the file manually.
func (s *SyncFile) Reload() error {
	if info, err := os.Stat(s.filename); err != nil {
		s.err = err
		return err
	} else if contents, err := ioutil.ReadFile(s.filename); err != nil {
		s.err = err
		return err
	} else if contents != nil {
		s.contents = contents
		s.reloadFunc(s)
		s.modTime = info.ModTime()
	}
	return nil
}

// Apply changes to the file.
func (s *SyncFile) Apply() error {
	if err := ioutil.WriteFile(s.filename, s.contents, os.ModePerm); err != nil {
		s.err = err
		return err
	}
	return nil
}

// Change the contents
func (s *SyncFile) Change(new []byte) error {
	s.contents = new
	return s.Apply()
}

// Stop goroutines and close channels.
func (s *SyncFile) Exit(wait bool) {
	s.control <- cmdExit
	if wait {
		s.wg.Wait()
	}
}

// Create a SyncFile object.
func NewSyncFile(filename string, interval time.Duration) (*SyncFile, error) {
	syncFile := &SyncFile{
		filename:   filename,
		interval:   interval,
		control:    make(chan command, 32),
		reloadFunc: func(i *SyncFile) {},
	}
	// Open file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Load File
	if _, err = file.Read(syncFile.contents); err != nil {
		return nil, err
	}
	// Get Info
	info, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	syncFile.modTime = info.ModTime()
	// Start loop-goroutine
	syncFile.wg.Add(1)
	go syncFile.loopGoroutine()
	return syncFile, nil
}
