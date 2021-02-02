# go语言用的常用开发工具包
关于每个功能的详细，请看源代码内的注释。
如果有问题和建议欢迎issue和pull-request
# Common utilities for golang.
About the details of functions, Please see the documentation in source code.
Welcome to issue & pull-request
# go言語用汎用ユーティリティライブラリ。
各関数の機能の詳細はソースコード内のコメントを確認してください。
issue & pull-request 大歓迎

# 安装，install, インストール
`go get github.com/el-ideal-ideas/ellib`

# 目录，index, 目次
- [elarr](#elarr) utilities for array & slice.
- [elconv](#elconv) utilities for data convert.
- [eldev](#eldev) utilities for debug.
- [eldump](#eldump) utilities for dump data to debug.
- [elenv](#elenv) utilities for system environment.
- [elfs](#elfs) utilities for filesystem.
- [elinfo](#elinfo) information about this package.
- [eljson](#eljson) wrapper of github.com/json-iterator/go
- [elmap](#elmap) utilities for map.
- [elmath](#elmath) utilities for math.
- [elnet](#elnet) utilities for network
- [elrand](#elrand) utilities for random.
- [elref](#elref) utilities used reflect.
- [elstr](#elstr) utilities for string.
- [elsyncfile](#elsyncfile) this package can sync struct with file.
- [elsys](#elsys) utilities for system.
- [eltengo](#eltengo) utilities for tengo script.
- [elutils](#elutils) some common utilities.

## elarr
- 数组 & 切片
- array & slice
- 配列&スライス関連機能
```go
import "github.com/el-ideal-ideas/ellib/elarr"

// 类型转换 (interface{})
// Convert (interface{})
// 型変換系 (interface{})
elarr.ToStrings(arr []interface{}) []string
elarr.ToInts(arr []interface{}) []int
elarr.ToInt8s(arr []interface{}) []int8
elarr.ToInt16s(arr []interface{}) []int16
elarr.ToInt32s(arr []interface{}) []int32
elarr.ToInt64s(arr []interface{}) []int64
elarr.ToUints(arr []interface{}) []uint
elarr.ToUint8s(arr []interface{}) []uint8
elarr.ToUint16s(arr []interface{}) []uint16
elarr.ToUint32s(arr []interface{}) []uint32
elarr.ToUint64s(arr []interface{}) []uint64
elarr.ToBytes(arr []interface{}) []byte
elarr.ToRunes(arr []interface{}) []rune
elarr.ToFloat32s(arr []interface{}) []float32
elarr.ToFloat64s(arr []interface{}) []float64
// 类型转换 (转换到字符串的切片)
// Convert (to []string)
// 型変換系 ([]string型に変換)
elarr.IntsToStrings(arr []int) []string
elarr.Int8sToStrings(arr []int8) []string
elarr.Int16sToStrings(arr []int16) []string
elarr.Int32sToStrings(arr []int32) []string
elarr.Int64sToStrings(arr []int64) []string
elarr.UintsToStrings(arr []uint) []string
elarr.Uint8sToStrings(arr []uint8) []string
elarr.Uint16sToStrings(arr []uint16) []string
elarr.Uint32sToStrings(arr []uint32) []string
elarr.Uint64sToStrings(arr []uint64) []string
elarr.RunesToStrings(arr []rune) []string
elarr.BytesToStrings(arr []byte) []string
elarr.Float32sToStrings(arr []float32, fmt byte, prec, bitSize int) []string
elarr.Float64sToStrings(arr []float64, fmt byte, prec, bitSize int) []string
// 类型转换 (其他)
// Convert (Other)
// 型変換系 (その他)
elarr.IntsToFloat64s(arr []int) []float64
elarr.Float64sToInts(arr []float64) []int
elarr.IntsToUints(arr []int) []uint
elarr.UintsToInts(arr []uint) []int
elarr.UintsToFloat64s(arr []uint) []float64
elarr.Float64sToUints(arr []float64) []uint
elarr.StringsToInts(arr []string) ([]int, error)
elarr.StringsToUints(arr []string) ([]uint, error)
elarr.StringsToFloat64s(arr []string) ([]float64, error)
elarr.RunesToBytes(rs []rune) []byte

// ----------------------------------------------------------------------------

// 用指定数据填充切片
// Fill slice with target data
// 指定のデータでスライスを満たす
elarr.FillInter(arr []interface{}, item interface{})
elarr.FillStr(arr []string, item string)
elarr.FillInt(arr []int, item int)
elarr.FillInt8(arr []int8, item int8)
elarr.FillInt16(arr []int16, item int16)
elarr.FillInt32(arr []int32, item int32)
elarr.FillInt64(arr []int64, item int64)
elarr.FillUint(arr []uint, item uint)
elarr.FillUint8(arr []uint8, item uint8)
elarr.FillUint16(arr []uint16, item uint16)
elarr.FillUint32(arr []uint32, item uint32)
elarr.FillUint64(arr []uint64, item uint64)
elarr.FillRune(arr []rune, item rune)
elarr.FillByte(arr []byte, item byte)
elarr.FillFloat32(arr []float32, item float32)
elarr.FillFloat64(arr []float64, item float64)

// ----------------------------------------------------------------------------

// 类似于其他语言的foreach或者for-in的功能
// foreach (likes other languages)
// 他の言語でforeachまたはfor-inなどと呼ばれる機能
elarr.ForEachInter(arr []interface{}, f func(interface{}))
elarr.ForEachStr(arr []string, f func(string))
elarr.ForEachInt(arr []int, f func(int))
elarr.ForEachInt8(arr []int8, f func(int8))
elarr.ForEachInt16(arr []int16, f func(int16))
elarr.ForEachInt32(arr []int32, f func(int32))
elarr.ForEachInt64(arr []int64, f func(int64))
elarr.ForEachUint(arr []uint, f func(uint))
elarr.ForEachUint8(arr []uint8, f func(uint8))
elarr.ForEachUint16(arr []uint16, f func(uint16))
elarr.ForEachUint32(arr []uint32, f func(uint32))
elarr.ForEachUint64(arr []uint64, f func(uint64))
elarr.ForEachRune(arr []rune, f func(rune))
elarr.ForEachByte(arr []byte, f func(byte))
elarr.ForEachFloat32(arr []float32, f func(float32))
elarr.ForEachFloat64(arr []float64, f func(float64))

// ----------------------------------------------------------------------------

// 如果数组中包含指定的值，则返回true
// If the slice contains target value, return true
// スライスが指定のデータを含んで入れば、trueを返す
elarr.InF(data []interface{}, f func(interface{}) bool) bool
elarr.InInter(data []interface{}, item interface{}) bool
elarr.InStr(data []string, item string) bool
elarr.InInt(data []int, item int) bool
elarr.InInt8(data []int8, item int8) bool 
elarr.InInt16(data []int16, item int16) bool
elarr.InInt32(data []int32, item int32) bool
elarr.InInt64(data []int64, item int64) bool
elarr.InUint(data []uint, item uint) bool
elarr.InUint8(data []uint8, item uint8) bool
elarr.InUint16(data []uint16, item uint16) bool
elarr.InUint32(data []uint32, item uint32) bool
elarr.InUint64(data []uint64, item uint64) bool
elarr.InRune(data []rune, item rune)
elarr.InByte(data []byte, item byte) bool
elarr.InFloat32(data []float32, item float32) bool
elarr.InFloat64(data []float64, item float64) bool

// ----------------------------------------------------------------------------

// 如果参数内的所有数组内容相等，则返回true
// If all slice in args, return true
// 引数として渡されたすべてのスライスが同等である場合, trueを返す
elarr.IsSame(data ...interface{}) bool
elarr.IsSameInter(data ...interface{}) bool
elarr.IsSameInterSlice(data ...[]interface{}) bool
elarr.IsSameStr(data ...[]string) bool
elarr.IsSameInt(data ...[]int) bool
elarr.IsSameInt8(data ...[]int8) bool
elarr.IsSameInt16(data ...[]int16) bool
elarr.IsSameInt32(data ...[]int32) bool
elarr.IsSameInt64(data ...[]int64) bool
elarr.IsSameUint(data ...[]uint) bool
elarr.IsSameUint8(data ...[]uint8) bool
elarr.IsSameUint16(data ...[]uint16) bool
elarr.IsSameUint32(data ...[]uint32) bool
elarr.IsSameUint64(data ...[]uint64) bool
elarr.IsSameRune(data ...[]rune) bool
elarr.IsSameByte(data ...[]byte) bool
elarr.IsSameFloat32(data ...[]float32) bool
elarr.IsSameFloat64(data ...[]float64) bool

// ----------------------------------------------------------------------------

// 类似于python的map函数。
// Like python's map function.
// pythonのmap関数に類似した機能を提供
elarr.Map(arr []interface{}, f func(interface{}) interface{}) []interface{}
elarr.MapStrInter(arr []string, f func(string) interface{}) []interface{}
elarr.MapStrStr(arr []string, f func(string) string) []string
elarr.MapStrInt(arr []string, f func(string) int) []int
elarr.MapStrUint(arr []string, f func(string) uint) []uint
elarr.MapStrFloat64(arr []string, f func(string) float64) []float64
elarr.MapIntInter(arr []int, f func(int) interface{}) []interface{}
elarr.MapIntStr(arr []int, f func(int) string) []string
elarr.MapIntInt(arr []int, f func(int) int) []int
elarr.MapIntUint(arr []int, f func(int) uint) []uint
elarr.MapIntFloat64(arr []int, f func(int) float64) []float64
elarr.MapFloat64Inter(arr []float64, f func(float64) interface{}) []interface{}
elarr.MapFloat64Str(arr []float64, f func(float64) string) []string
elarr.MapFloat64Int(arr []float64, f func(float64) int) []int
elarr.MapFloat64Uint(arr []float64, f func(float64) uint) []uint
elarr.MapFloat64Float64(arr []float64, f func(float64) float64) []float64
elarr.MapUintInter(arr []uint, f func(uint) interface{}) []interface{}
elarr.MapUintStr(arr []uint, f func(uint) string) []string
elarr.MapUintInt(arr []uint, f func(uint) int) []int
elarr.MapUintUint(arr []uint, f func(uint) uint) []uint
elarr.MapUintFloat64(arr []uint, f func(uint) float64) []float64
elarr.MapInterInter(arr []interface{}, f func(interface{}) interface{}) []interface{}
elarr.MapInterStr(arr []interface{}, f func(interface{}) string) []string
elarr.MapInterInt(arr []interface{}, f func(interface{}) int) []int
elarr.MapInterUint(arr []interface{}, f func(interface{}) uint) []uint
elarr.MapInterFloat64(arr []interface{}, f func(interface{}) float64) []float64

// ----------------------------------------------------------------------------

// 把指定的数组或者切片反过来
// reverse target array or slice
// 配列＆スライスを反転させる
elarr.Reverse(s []interface{})
elarr.ReverseStr(s []string)
elarr.ReverseUint(s []uint)
elarr.ReverseUint8(s []uint8)
elarr.ReverseUint16(s []uint16)
elarr.ReverseUint32(s []uint32)
elarr.ReverseUint64(s []uint64)
elarr.ReverseInt(s []int)
elarr.ReverseInt8(s []int8)
elarr.ReverseInt16(s []int16)
elarr.ReverseInt32(s []int32)
elarr.ReverseInt64(s []int64)
elarr.ReverseByte(s []byte)
elarr.ReverseRune(s []rune)
elarr.ReverseFloat32(s []float32)
elarr.ReverseFloat64(s []float64)

// ----------------------------------------------------------------------------

// 消除重复数据
// Remove duplicates item
// 重複アイテムを消す
elarr.UniqueInter(arr []interface{}) []interface{}
elarr.UniqueStr(arr []string) []string
elarr.UniqueInt(arr []int) []int
elarr.UniqueInt8(arr []int8) []int8
elarr.UniqueInt16(arr []int16) []int16
elarr.UniqueInt32(arr []int32) []int32
elarr.UniqueInt64(arr []int64) []int64
elarr.UniqueUint(arr []uint) []uint
elarr.UniqueUint8(arr []uint8) []uint8
elarr.UniqueUint16(arr []uint16) []uint16
elarr.UniqueUint32(arr []uint32) []uint32
elarr.UniqueUint64(arr []uint64) []uint64
elarr.UniqueByte(arr []byte) []byte
elarr.UniqueRune(arr []rune) []rune
elarr.UniqueFloat32(arr []float32) []float32
elarr.UniqueFloat64(arr []float64) []float64

// ----------------------------------------------------------------------------

// 消除不需要的数据
// filter items in slice or array
// 不要なデータを削除
elarr.Filter(data interface{}, condition func(interface{})bool) []interface{}
elarr.FilterInter(data []interface{}, condition func(interface{})bool) []interface{}
elarr.FilterStr(data []string, condition func(string)bool) []string
elarr.FilterInt(data []int, condition func(int)bool) []int
elarr.FilterInt8(data []int8, condition func(int8)bool) []int8
elarr.FilterInt16(data []int16, condition func(int16)bool) []int16
elarr.FilterInt32(data []int32, condition func(int32)bool) []int32
elarr.FilterInt64(data []int64, condition func(int64)bool) []int64
elarr.FilterUint(data []uint, condition func(uint)bool) []uint
elarr.FilterUint8(data []uint8, condition func(uint8)bool) []uint8
elarr.FilterUint16(data []uint16, condition func(uint16)bool) []uint16
elarr.FilterUint32(data []uint32, condition func(uint32)bool) []uint32
elarr.FilterUint64(data []uint64, condition func(uint64)bool) []uint64
elarr.FilterRune(data []rune, condition func(rune)bool) []rune
elarr.FilterByte(data []byte, condition func(byte)bool) []byte
elarr.FilterFloat32(data []float32, condition func(float32)bool) []float32
elarr.FilterFloat64(data []float64, condition func(float64)bool) []float64

```

## elconv
- interface{}转换成指定类型
- convert interface{} to other type
- interface{}を特定の型に変換する
```go
import "github.com/el-ideal-ideas/ellib/elconv"

elconv.AsBool(v interface{}) bool
elconv.AsFloat64(v interface{}) float64
elconv.AsFloat32(v interface{}) float32
elconv.AsInt(v interface{}) int
elconv.AsInt64(v interface{}) int64
elconv.AsStr(v interface{}) string
elconv.AsUint(v interface{}) uint
elconv.AsUint64(v interface{}) uint64
```

## elfs
- 用于处理文件系统的包
- utilities for filesystem
- ファイルシステム関連のユーティリティ
```go
import "github.com/el-ideal-ideas/ellib/elfs"

elfs.PathExists(path string) bool
elfs.FileExists(path string) bool
elfs.IsDir(path string) bool
elfs.IsFile(path string) bool
elfs.MimeType(path string) (mime string)
elfs.ReaderMimeType(r io.Reader) (mime string)
elfs.IsImageFile(path string) bool
elfs.DeleteIfFileExist(filename string) error
elfs.SelfDir() (path string, err error)
elfs.GetTempDir() string
elfs.ParentDir(path string) (string, error)
elfs.GetFileExt(filename string) string
elfs.Home() string
elfs.SetHome(s string) error

// elfs 包含了 path/filepath 的全部函数
// elfs contain alias of path/filepath
// elfs は path/filepath　にあるすべての関数を含む
var Join = filepath.Join
var Abs = filepath.Abs
var Split = filepath.Split
var Match = filepath.Match
var Base = filepath.Base
var Clean = filepath.Clean
var Dir = filepath.Dir
var ErrBadPattern = filepath.ErrBadPattern
var EvalSymlinks = filepath.EvalSymlinks
var Ext = filepath.Ext
var FromSlash = filepath.FromSlash
var Glob = filepath.Glob
var IsAbs = filepath.IsAbs
var ListSeparator = filepath.ListSeparator
var Rel = filepath.Rel
var Separator = filepath.Separator
var SkipDir = filepath.SkipDir
var SplitList = filepath.SplitList
var ToSlash = filepath.ToSlash
var VolumeName = filepath.VolumeName
var Walk = filepath.Walk
type WalkFunc = filepath.WalkFunc
```

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

## eldump
- 这个包可以把变量内容用容易理解格式打印到标准输出，方便开发和测试。
  这个包是从github.com/gookit/goutil复制过来的，目前对代码并无更改。
- This package can print data more clear and beautiful, For debug & development.
  This package was copy from github.com/gookit/goutil.
- 開発＆デバッグ用に変数の中身を見やすい形式で出力する。
  このパッケージはgithub.com/gookit/goutilからコピーされたものであり、現時点では変更点はありません。
```go
package main

import "github.com/el-ideal-ideas/ellib/eldump"

func main() {
	eldump.V(map[string]interface{}{
		"key00": []int{1, 2, 3, 4},
		"key01": map[int]int{
			1: 2,
			2: 4,
			3: 6,
		},
		"key02": "some text",
	})
}
```
```
PRINT AT main.main(main.go:6)
map[string]interface {} {
  "key00": []int{1, 2, 3, 4},
  "key01": map[int]int {
    2: int(4),
    3: int(6),
    1: int(2),
  },
  "key02": string("some text"),
},
```

## elenv
- 关于环境变量的包
- utilities for system environments
- 環境変数関連のユーティリティ
```go
import "github.com/el-ideal-ideas/ellib/elenv"

elenv.GetEnv(name string, def ...string) string
elenv.Setenv(key, value string) error
```

## eljson
- 用 github.com/json-iterator/go 代替原生json包
- Use github.com/json-iterator/go instead of default json package 
- デフォルトのjsonパッケージの代わりにgithub.com/json-iterator/goを使用するように、ラップしたもの
```go
import json "github.com/el-ideal-ideas/ellib/eljson"

var Get = json.Get
var Marshal = json.Marshal
var MarshalIndent = json.MarshalIndent
var BorrowIterator = json.BorrowIterator
var BorrowStream = json.BorrowStream
var DecoderOf = json.DecoderOf
var EncoderOf = json.EncoderOf
var MarshalToString = json.MarshalToString
var NewDecoder = json.NewDecoder
var NewEncoder = json.NewEncoder
var RegisterExtension = json.RegisterExtension
var Valid = json.Valid
var UnmarshalFromString = json.UnmarshalFromString
var ReturnStream = json.ReturnStream
var Unmarshal = json.Unmarshal
var ReturnIterator = json.ReturnIterator
```

## elinfo
- 关于ellib包的信息
- Info about ellib
- ellibパッケージに関する情報
```go
const Version = "0.1.0"
```
