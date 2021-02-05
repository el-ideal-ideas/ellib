#### elstr包含官方strings包的全部函数 (为了性能优化一部分函数用了其他的实现方式)
#### elstr contains all functions in official strings package (some functions use different source code for performance)
#### elstrパブリックは公式のstringsパッケージの全部の関数を含んでいます (一部の関数は性能向上のため、異なる実装になっています)

```go
type Builder strings.Builder
type Reader strings.Reader
type Replacer strings.Replacer

var Compare = strings.Compare
var NewReader = strings.NewReader
var NewReplacer = strings.NewReplacer
var SCount = strings.Count
var Contains = strings.Contains
var ContainsAny = strings.ContainsAny
var ContainsRune = strings.ContainsRune
var LastIndex = strings.LastIndex
var IndexByte = strings.IndexByte
var IndexRune = strings.IndexRune
var IndexAny = strings.IndexAny
var LastIndexAny = strings.LastIndexAny
var LastIndexByte = strings.LastIndexByte
var SplitN = strings.SplitN
var SplitAfterN = strings.SplitAfterN
var Split = strings.Split
var SplitAfter = strings.SplitAfter
var Fields = strings.Fields
var FieldsFunc = strings.FieldsFunc
var SJoin = strings.Join
var HasPrefix = strings.HasPrefix
var HasSuffix = strings.HasSuffix
var Map = strings.Map
var SRepeat = strings.Repeat
var ToUpper = strings.ToUpper
var ToLower = strings.ToLower
var ToTitle = strings.ToTitle
var ToUpperSpecial = strings.ToUpperSpecial
var ToLowerSpecial = strings.ToLowerSpecial
var ToTitleSpecial = strings.ToTitleSpecial
var ToValidUTF8 = strings.ToValidUTF8
var Title = strings.Title
var TrimLeftFunc = strings.TrimLeftFunc
var TrimRightFunc = strings.TrimRightFunc
var TrimFunc = strings.TrimFunc
var IndexFunc = strings.IndexFunc
var LastIndexFunc = strings.LastIndexFunc
var Trim = strings.Trim
var TrimLeft = strings.TrimLeft
var TrimRight = strings.TrimRight
var TrimSpace = strings.TrimSpace
var TrimPrefix = strings.TrimPrefix
var TrimSuffix = strings.TrimSuffix
var SReplace = strings.Replace
var ReplaceAll = strings.ReplaceAll
var EqualFold = strings.EqualFold
var Index = strings.Index
```