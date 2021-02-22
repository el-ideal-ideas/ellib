package elsys

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"syscall"
)

// Exit system with additional information.
func Exit(code int, info string, err error) {
	fmt.Printf("Info: %s\n", info)
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("ExitCode: %d\n", code)
	syscall.Exit(code)
}

// run garbage collection manually (blocking).
func ForceGC() {
	runtime.GC()
	debug.FreeOSMemory()
}

// run garbage collection manually (non-blocking).
func ForceGCNonBlocking() {
	go func() {
		ForceGC()
	}()
}
