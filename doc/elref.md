## elref
- 关于reflect的工具包
- utilities for reflect
- reflect関連のユーティリティ
```go
import "github.com/el-ideal-ideas/ellib/elref"

// 检查是否为nil
// if `v` is nil return true
// nilチェック
elref.IsNil(v interface{}) bool

// 检查是否为空值
// if f `v` is empty value, return true
// 空の値(初期値)かどうかをチェック
elref.IsEmpty(v interface{}) bool

// 对array，slice，map，检测对象数据（或者是map的key）是否存在
// if `needle` is in `haystack` return true. supports array, slice and map.
// array, slice, map に指定の要素またはキーが存在するかを確認する
elref.IsIn(needle interface{}, haystack interface{}) bool

// 获取对象类型
// get the type of target value
// 指定の値のデータ型を取得
elref.Type(v interface{}) string

// 检测对象类型
// check the data type
// 型チェック
elref.IsStruct(v interface{}) bool
elref.IsInterface(v interface{}) bool
elref.IsMap(v interface{}) bool
elref.IsArray(v interface{}) bool
elref.IsSlice(v interface{}) bool
elref.IsArrayOrSlice(v interface{}) bool
elref.IsBool(v interface{}) bool
elref.IsInt(v interface{}) bool
elref.IsInt8(v interface{}) bool
elref.IsInt16(v interface{}) bool
elref.IsInt32(v interface{}) bool
elref.IsInt64(v interface{}) bool
elref.IsUint(v interface{}) bool
elref.IsUint8(v interface{}) bool
elref.IsUint16(v interface{}) bool
elref.IsUint32(v interface{}) bool
elref.IsUint64(v interface{}) bool
elref.IsUintptr(v interface{}) bool
elref.IsFloat32(v interface{}) bool
elref.IsFloat64(v interface{}) bool
elref.IsComplex64(v interface{}) bool
elref.IsComplex128(v interface{}) bool
elref.IsChan(v interface{}) bool
elref.IsFunc(v interface{}) bool
elref.IsPtr(v interface{}) bool
elref.IsString(v interface{}) bool
elref.IsUnsafePointer(v interface{})

```