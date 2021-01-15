package sys

import (
	"fmt"
	"syscall"
)


// Exit system with additional information.
func Exit(code int, info string, err error) {
	fmt.Printf("Info: %s, Error: %v\n", info, err)
	fmt.Printf("Exit<%d>", code)
	syscall.Exit(code)
}
