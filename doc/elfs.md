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