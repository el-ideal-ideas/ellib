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