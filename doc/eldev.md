## eldev
- 用于方便debug的函数群
- functions for debug
- デバッグ用パッケージ
```go
import "github.com/el-ideal-ideas/ellib/eldev"

eldev.GetFuncName(f interface{}) string
eldev.GetCurrentLine() int
eldev.GetCurrentFile() string
eldev.GetCallerLine() int
eldev.GetCallerFile() string
eldev.GetCallerName() string
eldev.GetPkgName() string
eldev.PanicIfErr(err error)
eldev.DumpStacks(depth ...uint)
eldev.GetCallStacks(all bool) []byte
eldev.GetType(v interface{}) string
eldev.Timeit(f func(), number int) time.Duration
eldev.RunSurveillance(d time.Duration, f func(*Surveillance))
```
#### About eldev.RunSurveillance
- 这个函数会启动一个goroutine并且开始实时监视系统状况，包含了最新的系统数据的Surveillance类型的变量，
  会被作为参数定期传给函数`f`。
- This function will Start a goroutine and check system stats with an interval.
  function `f` will be got a Surveillance object that contains the latest system info.
- この関数は新しいgoroutineを立ち上げ、システムの状態を監視します。
  システムの最新情報を含んだSurveillanceのオブジェクトは定期的に引数として、
  関数`f`に渡されます。
```go
type Surveillance struct {
	Time         int64
	TimeForHuman string
	NumGoroutine int
	MemStats     *runtime.MemStats
}
```
