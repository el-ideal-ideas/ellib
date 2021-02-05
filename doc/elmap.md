## elmap
- 关于map类型的各种函数
- functions for map type data.
- map型用ユーティリティ群
```go
import json "github.com/el-ideal-ideas/ellib/elmap"

// 合并2个map类型的数据
// merge src to dst map
// srcのデータをdstに統合する
elmap.MergeSSMap(src, dst map[string]string, ignoreCase bool)
elmap.MergeSIMap(src, dst map[string]int, ignoreCase bool)
elmap.MergeSUMap(src, dst map[string]uint, ignoreCase bool)
elmap.MergeSFMap(src, dst map[string]float64, ignoreCase bool)
elmap.MergeIIMap(src, dst map[int]int)
elmap.MergeIUMap(src, dst map[int]uint)
elmap.MergeISMap(src, dst map[int]string)
elmap.MergeIFMap(src, dst map[int]float64)
elmap.MergeUIMap(src, dst map[uint]int)
elmap.MergeUUMap(src, dst map[uint]uint)
elmap.MergeUSMap(src, dst map[uint]string)
elmap.MergeUFMap(src, dst map[uint]float64)

// 互换map的key和value
// flip key and value in the map
// mapのキーと値を交換する
elmap.MapFlip(m interface{}) map[interface{}]interface{}

// 获取map的全部值
// get all values in the map
// mapの値を全部取得する
elmap.Values(m interface{}) []interface{}

// 把所有的key转换成小写
// convert keys to lower
// 全てのキーを小文字に変換
elmap.KeysToLower(m map[string]interface{}) map[string]interface{}

// 把所有的key转换成大写
// convert keys to upper
// 全てのキーを大文字に変換
elmap.KeysToUpper(m map[string]interface{}) map[string]interface{}

// 获取map的全部key
// get all keys in the map
// map内の全てのキーを取得
elmap.Keys(m interface{}) []interface{}
elmap.KeysAsString(m interface{}) []string
elmap.KeysInt(m map[int]int) []int
elmap.KeysIntInt64(m map[int]int64) []int
elmap.KeysIntInt32(m map[int]int32) []int
elmap.KeysIntInt16(m map[int]int16) []int
elmap.KeysIntInt8(m map[int]int8) []int
elmap.KeysIntUint(m map[int]uint) []int
elmap.KeysIntUint64(m map[int]uint64) []int
elmap.KeysIntUint32(m map[int]uint32) []int
elmap.KeysIntUint16(m map[int]uint16) []int
elmap.KeysIntUint8(m map[int]uint8) []int
elmap.KeysIntStr(m map[int]string) []int
elmap.KeysIntBool(m map[int]bool) []int
elmap.KeysIntFloat32(m map[int]float32) []int
elmap.KeysIntFloat64(m map[int]float64) []int
elmap.KeysInt64Int(m map[int64]int) []int64
elmap.KeysInt64(m map[int64]int64) []int64
elmap.KeysInt64Int32(m map[int64]int32) []int64
elmap.KeysInt64Int16(m map[int64]int16) []int64
elmap.KeysInt64Int8(m map[int64]int8) []int64
elmap.KeysInt64Uint(m map[int64]uint) []int64
elmap.KeysInt64Uint64(m map[int64]uint64) []int64
elmap.KeysInt64Uint32(m map[int64]uint32) []int64
elmap.KeysInt64Uint16(m map[int64]uint16) []int64
elmap.KeysInt64Uint8(m map[int64]uint8) []int64
elmap.KeysInt64Str(m map[int64]string) []int64
elmap.KeysInt64Bool(m map[int64]bool) []int64
elmap.KeysInt64Float32(m map[int64]float32) []int64
elmap.KeysInt64Float64(m map[int64]float64) []int64
elmap.KeysInt32Int(m map[int32]int) []int32
elmap.KeysInt32Int64(m map[int32]int64) []int32
elmap.KeysInt32(m map[int32]int32) []int32
elmap.KeysInt32Int16(m map[int32]int16) []int32
elmap.KeysInt32Int8(m map[int32]int8) []int32
elmap.KeysInt32Uint(m map[int32]uint) []int32
elmap.KeysInt32Uint64(m map[int32]uint64) []int32
elmap.KeysInt32Uint32(m map[int32]uint32) []int32
elmap.KeysInt32Uint16(m map[int32]uint16) []int32
elmap.KeysInt32Uint8(m map[int32]uint8) []int32
elmap.KeysInt32Str(m map[int32]string) []int32
elmap.KeysInt32Bool(m map[int32]bool) []int32
elmap.KeysInt32Float32(m map[int32]float32) []int32
elmap.KeysInt32Float64(m map[int32]float64) []int32
elmap.KeysInt16Int(m map[int16]int) []int16
elmap.KeysInt16Int64(m map[int16]int64) []int16
elmap.KeysInt16Int32(m map[int16]int32) []int16
elmap.KeysInt16(m map[int16]int16) []int16
elmap.KeysInt16Int8(m map[int16]int8) []int16
elmap.KeysInt16Uint(m map[int16]uint) []int16
elmap.KeysInt16Uint64(m map[int16]uint64) []int16
elmap.KeysInt16Uint32(m map[int16]uint32) []int16
elmap.KeysInt16Uint16(m map[int16]uint16) []int16
elmap.KeysInt16Uint8(m map[int16]uint8) []int16
elmap.KeysInt16Str(m map[int16]string) []int16
elmap.KeysInt16Bool(m map[int16]bool) []int16
elmap.KeysInt16Float32(m map[int16]float32) []int16
elmap.KeysInt16Float64(m map[int16]float64) []int16
elmap.KeysInt8Int(m map[int8]int) []int8
elmap.KeysInt8Int64(m map[int8]int64) []int8
elmap.KeysInt8Int32(m map[int8]int32) []int8
elmap.KeysInt8Int16(m map[int8]int16) []int8
elmap.KeysInt8(m map[int8]int8) []int8
elmap.KeysInt8Uint(m map[int8]uint) []int8
elmap.KeysInt8Uint64(m map[int8]uint64) []int8
elmap.KeysInt8Uint32(m map[int8]uint32) []int8
elmap.KeysInt8Uint16(m map[int8]uint16) []int8
elmap.KeysInt8Uint8(m map[int8]uint8) []int8
elmap.KeysInt8Str(m map[int8]string) []int8
elmap.KeysInt8Bool(m map[int8]bool) []int8
elmap.KeysInt8Float32(m map[int8]float32) []int8
elmap.KeysInt8Float64(m map[int8]float64) []int8
elmap.KeysUintInt(m map[uint]int) []uint
elmap.KeysUintInt64(m map[uint]int64) []uint
elmap.KeysUintInt32(m map[uint]int32) []uint
elmap.KeysUintInt16(m map[uint]int16) []uint
elmap.KeysUintInt8(m map[uint]int8) []uint
elmap.KeysUint(m map[uint]uint) []uint
elmap.KeysUintUint64(m map[uint]uint64) []uint
elmap.KeysUintUint32(m map[uint]uint32) []uint
elmap.KeysUintUint16(m map[uint]uint16) []uint
elmap.KeysUintUint8(m map[uint]uint8) []uint
elmap.KeysUintStr(m map[uint]string) []uint
elmap.KeysUintBool(m map[uint]bool) []uint
elmap.KeysUintFloat32(m map[uint]float32) []uint
elmap.KeysUintFloat64(m map[uint]float64) []uint
elmap.KeysUint64Int(m map[uint64]int) []uint64
elmap.KeysUint64Int64(m map[uint64]int64) []uint64
elmap.KeysUint64Int32(m map[uint64]int32) []uint64
elmap.KeysUint64Int16(m map[uint64]int16) []uint64
elmap.KeysUint64Int8(m map[uint64]int8) []uint64
elmap.KeysUint64Uint(m map[uint64]uint) []uint64
elmap.KeysUint64(m map[uint64]uint64) []uint64
elmap.KeysUint64Uint32(m map[uint64]uint32) []uint64
elmap.KeysUint64Uint16(m map[uint64]uint16) []uint64
elmap.KeysUint64Uint8(m map[uint64]uint8) []uint64
elmap.KeysUint64Str(m map[uint64]string) []uint64
elmap.KeysUint64Bool(m map[uint64]bool) []uint64
elmap.KeysUint64Float32(m map[uint64]float32) []uint64
elmap.KeysUint64Float64(m map[uint64]float64) []uint64
elmap.KeysUint32Int(m map[uint32]int) []uint32
elmap.KeysUint32Int64(m map[uint32]int64) []uint32
elmap.KeysUint32Int32(m map[uint32]int32) []uint32
elmap.KeysUint32Int16(m map[uint32]int16) []uint32
elmap.KeysUint32Int8(m map[uint32]int8) []uint32
elmap.KeysUint32Uint(m map[uint32]uint) []uint32
elmap.KeysUint32Uint64(m map[uint32]uint64) []uint32
elmap.KeysUint32(m map[uint32]uint32) []uint32
elmap.KeysUint32Uint16(m map[uint32]uint16) []uint32
elmap.KeysUint32Uint8(m map[uint32]uint8) []uint32
elmap.KeysUint32Str(m map[uint32]string) []uint32
elmap.KeysUint32Bool(m map[uint32]bool) []uint32
elmap.KeysUint32Float32(m map[uint32]float32) []uint32
elmap.KeysUint32Float64(m map[uint32]float64) []uint32
elmap.KeysUint16Int(m map[uint16]int) []uint16
elmap.KeysUint16Int64(m map[uint16]int64) []uint16
elmap.KeysUint16Int32(m map[uint16]int32) []uint16
elmap.KeysUint16Int16(m map[uint16]int16) []uint16
elmap.KeysUint16Int8(m map[uint16]int8) []uint16
elmap.KeysUint16Uint(m map[uint16]uint) []uint16
elmap.KeysUint16Uint64(m map[uint16]uint64) []uint16
elmap.KeysUint16Uint32(m map[uint16]uint32) []uint16
elmap.KeysUint16(m map[uint16]uint16) []uint16
elmap.KeysUint16Uint8(m map[uint16]uint8) []uint16
elmap.KeysUint16Str(m map[uint16]string) []uint16
elmap.KeysUint16Bool(m map[uint16]bool) []uint16
elmap.KeysUint16Float32(m map[uint16]float32) []uint16
elmap.KeysUint16Float64(m map[uint16]float64) []uint16
elmap.KeysUint8Int(m map[uint8]int) []uint8
elmap.KeysUint8Int64(m map[uint8]int64) []uint8
elmap.KeysUint8Int32(m map[uint8]int32) []uint8
elmap.KeysUint8Int16(m map[uint8]int16) []uint8
elmap.KeysUint8Int8(m map[uint8]int8) []uint8
elmap.KeysUint8Uint(m map[uint8]uint) []uint8
elmap.KeysUint8Uint64(m map[uint8]uint64) []uint8
elmap.KeysUint8Uint32(m map[uint8]uint32) []uint8
elmap.KeysUint8Uint16(m map[uint8]uint16) []uint8
elmap.KeysUint8(m map[uint8]uint8) []uint8
elmap.KeysUint8Str(m map[uint8]string) []uint8
elmap.KeysUint8Bool(m map[uint8]bool) []uint8
elmap.KeysUint8Float32(m map[uint8]float32) []uint8
elmap.KeysUint8Float64(m map[uint8]float64) []uint8
elmap.KeysStrInt(m map[string]int) []string
elmap.KeysStrInt64(m map[string]int64) []string
elmap.KeysStrInt32(m map[string]int32) []string
elmap.KeysStrInt16(m map[string]int16) []string
elmap.KeysStrInt8(m map[string]int8) []string
elmap.KeysStrUint(m map[string]uint) []string
elmap.KeysStrUint64(m map[string]uint64) []string
elmap.KeysStrUint32(m map[string]uint32) []string
elmap.KeysStrUint16(m map[string]uint16) []string
elmap.KeysStrUint8(m map[string]uint8) []string
elmap.KeysStr(m map[string]string) []string
elmap.KeysStrBool(m map[string]bool) []string
elmap.KeysStrFloat32(m map[string]float32) []string
elmap.KeysStrFloat64(m map[string]float64) []string
elmap.KeysBoolInt(m map[bool]int) []bool
elmap.KeysBoolInt64(m map[bool]int64) []bool
elmap.KeysBoolInt32(m map[bool]int32) []bool
elmap.KeysBoolInt16(m map[bool]int16) []bool
elmap.KeysBoolInt8(m map[bool]int8) []bool
elmap.KeysBoolUint(m map[bool]uint) []bool
elmap.KeysBoolUint64(m map[bool]uint64) []bool
elmap.KeysBoolUint32(m map[bool]uint32) []bool
elmap.KeysBoolUint16(m map[bool]uint16) []bool
elmap.KeysBoolUint8(m map[bool]uint8) []bool
elmap.KeysBoolStr(m map[bool]string) []bool
elmap.KeysBool(m map[bool]bool) []bool
elmap.KeysBoolFloat32(m map[bool]float32) []bool
elmap.KeysBoolFloat64(m map[bool]float64) []bool
elmap.KeysFloat32Int(m map[float32]int) []float32
elmap.KeysFloat32Int64(m map[float32]int64) []float32
elmap.KeysFloat32Int32(m map[float32]int32) []float32
elmap.KeysFloat32Int16(m map[float32]int16) []float32
elmap.KeysFloat32Int8(m map[float32]int8) []float32
elmap.KeysFloat32Uint(m map[float32]uint) []float32
elmap.KeysFloat32Uint64(m map[float32]uint64) []float32
elmap.KeysFloat32Uint32(m map[float32]uint32) []float32
elmap.KeysFloat32Uint16(m map[float32]uint16) []float32
elmap.KeysFloat32Uint8(m map[float32]uint8) []float32
elmap.KeysFloat32Str(m map[float32]string) []float32
elmap.KeysFloat32Bool(m map[float32]bool) []float32
elmap.KeysFloat32(m map[float32]float32) []float32
elmap.KeysFloat32Float64(m map[float32]float64) []float32
elmap.KeysFloat64Int(m map[float64]int) []float64
elmap.KeysFloat64Int64(m map[float64]int64) []float64
elmap.KeysFloat64Int32(m map[float64]int32) []float64
elmap.KeysFloat64Int16(m map[float64]int16) []float64
elmap.KeysFloat64Int8(m map[float64]int8) []float64
elmap.KeysFloat64Uint(m map[float64]uint) []float64
elmap.KeysFloat64Uint64(m map[float64]uint64) []float64
elmap.KeysFloat64Uint32(m map[float64]uint32) []float64
elmap.KeysFloat64Uint16(m map[float64]uint16) []float64
elmap.KeysFloat64Uint8(m map[float64]uint8) []float64
elmap.KeysFloat64Str(m map[float64]string) []float64
elmap.KeysFloat64Bool(m map[float64]bool) []float64
elmap.KeysFloat64Float32(m map[float64]float32) []float64
elmap.KeysFloat64(m map[float64]float64) []float64 
```