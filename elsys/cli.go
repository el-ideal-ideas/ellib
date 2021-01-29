package elsys

import (
	"bufio"
	"github.com/gookit/color"
	"os"
	"strings"
)

// Input read one line from user input.
// Usage:
// 	in := Input()
// 	ans, _ := Input("your name?")
func Input(s ...string) (string, error) {
	if len(s) > 0 && len(s[0]) > 0 {
		color.Print(s[0])
	}
	reader := bufio.NewReader(os.Stdin)
	answer, _, err := reader.ReadLine()
	return strings.TrimSpace(string(answer)), err
}
