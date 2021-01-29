// +build !windows

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
	bs, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		return ""
	}
	fmt.Println() // new line
	return string(bs)
}
