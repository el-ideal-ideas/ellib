// +build linux

package elsys

import (
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

// Exists check process running by given pid
func Exists(pid int) bool {
	if _, err := os.Stat(filepath.Join("/proc", strconv.Itoa(pid))); err == nil {
		return true
	}
	return false
}

// Kill process by pid
func Kill(pid int, signal syscall.Signal) error {
	return syscall.Kill(pid, signal)
}
