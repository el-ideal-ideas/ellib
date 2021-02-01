/*
	Example:

	executor := eltengo.New("result := a + b")
	result, _ := executor(map[string]interface{}{
		"a": "もずく",
		"b": "美味しい",
	})
	fmt.Println(result.String()) // もずく美味しい
	result, _ = executor(map[string]interface{}{
		"a": 20,
		"b": 7,
	})
	fmt.Println(result.Int64()) // 27
*/

package eltengo

import (
	"context"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
	"sync"
)

type Executor func(map[string]interface{}) (result *tengo.Variable, err error)

// Create a executor.
func New(s string) Executor {
	script := tengo.NewScript([]byte(defaultScript + s))
	script.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	lock := sync.Mutex{}
	return func(args map[string]interface{}) (result *tengo.Variable, err error) {
		lock.Lock()
		defer lock.Unlock()
		for k := range args {
			_ = script.Add(k, args[k])
		}
		var compiled *tengo.Compiled
		compiled, err = script.RunContext(context.Background())
		if err != nil {
			return nil, err
		} else {
			return compiled.Get("result"), nil
		}
	}
}
