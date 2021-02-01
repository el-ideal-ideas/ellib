package eltengo

import (
	"context"
	"errors"
	"fmt"
	"github.com/d5/tengo/v2"
	"github.com/el-ideal-ideas/ellib/elarr"
)

var allowed = []int32{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '+', '-', '*', '/', '%', '(', ')', '.'}

// This function can calculate simple data.
// For example:
// Calc("1+(2*3)/5")
func Calc(s string) (float64, error) {
	for _, c := range s {
		if !elarr.InInt32(allowed, c) {
			return 0, errors.New("invalid argument")
		}
	}
	script := tengo.NewScript([]byte(fmt.Sprintf("result := %s", s)))
	if compiled, err := script.RunContext(context.Background()); err == nil {
		return compiled.Get("result").Float(), nil
	} else {
		return 0, err
	}
}
