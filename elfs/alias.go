// alias of path/filepath

package elfs

import "path/filepath"

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
