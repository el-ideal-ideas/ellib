// +build windows

package elsys

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

// InputPassword from console terminal
func InputPassword(s ...string) string {
	if len(s) > 0 {
		fmt.Print(s[0])
	} else {
		fmt.Print("Enter Password: ")
	}
	// on windows, must convert 'syscall.Stdin' to int
	bs, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return ""
	}
	fmt.Println() // new line
	return string(bs)
}
