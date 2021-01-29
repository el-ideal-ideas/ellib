// +build darwin

package elsys

import (
	"golang.org/x/sys/unix"
	"syscall"
)

// Exists check process running by given pid
func Exists(pid int) bool {
	// OS X does not have a proc filesystem.
	// Use kill -0 pid to judge if the process exists.
	err := unix.Kill(pid, 0)
	return err == nil
}

// Kill process by pid
func Kill(pid int, signal syscall.Signal) error {
	return syscall.Kill(pid, signal)
}
