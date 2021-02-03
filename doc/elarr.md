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

// ----------------------------------------------------------------------------

// 从切片获取指定数据，并且从切片删除
// get a item from slice, and remove it form slice
// スライスから指定の要素を取り出し、その後削除する
elarr.Pop(v []interface{}, index int) (res interface{}, newPtr []interface{})
elarr.PopInter(v []interface{}, index int) (res interface{}, newPtr []interface{})
elarr.PopStr(v []string, index int) (res string, newPtr []string)
elarr.PopInt(v []int, index int) (res int, newPtr []int)
elarr.PopInt8(v []int8, index int) (res int8, newPtr []int8)
elarr.PopInt16(v []int16, index int) (res int16, newPtr []int16)
elarr.PopInt32(v []int32, index int) (res int32, newPtr []int32)
elarr.PopInt64(v []int64, index int) (res int64, newPtr []int64)
elarr.PopUint(v []uint, index int) (res uint, newPtr []uint)
elarr.PopUint8(v []uint8, index int) (res uint8, newPtr []uint8)
elarr.PopUint16(v []uint16, index int) (res uint16, newPtr []uint16)
elarr.PopUint32(v []uint32, index int) (res uint32, newPtr []uint32)
elarr.PopUint64(v []uint64, index int) (res uint64, newPtr []uint64)
elarr.PopRune(v []rune, index int) (res rune, newPtr []rune)
elarr.PopByte(v []byte, index int) (res byte, newPtr []byte)
elarr.PopFloat32(v []float32, index int) (res float32, newPtr []float32)
elarr.PopFloat64(v []float64, index int) (res float64, newPtr []float64)

// ----------------------------------------------------------------------------

// 在切片的指定位置插入数据
// insert a item
// 指定位置に要素を挿入
elarr.Insert(v []interface{}, index int, value interface{}) (newPtr []interface{})
elarr.InsertInter(v []interface{}, index int, value interface{}) (newPtr []interface{})
elarr.InsertStr(v []string, index int, value string) (newPtr []string)
elarr.InsertInt(v []int, index int, value int) (newPtr []int)
elarr.InsertInt8(v []int8, index int, value int8) (newPtr []int8)
elarr.InsertInt16(v []int16, index int, value int16) (newPtr []int16)
elarr.InsertInt32(v []int32, index int, value int32) (newPtr []int32)
elarr.InsertInt64(v []int64, index int, value int64) (newPtr []int64)
elarr.InsertUint(v []uint, index int, value uint) (newPtr []uint)
elarr.InsertUint8(v []uint8, index int, value uint8) (newPtr []uint8)
elarr.InsertUint16(v []uint16, index int, value uint16) (newPtr []uint16)
elarr.InsertUint32(v []uint32, index int, value uint32) (newPtr []uint32)
elarr.InsertUint64(v []uint64, index int, value uint64) (newPtr []uint64)
elarr.InsertRune(v []rune, index int, value rune) (newPtr []rune)
elarr.InsertByte(v []byte, index int, value byte) (newPtr []byte)
elarr.InsertFloat32(v []float32, index int, value float32) (newPtr []float32)
elarr.InsertFloat64(v []float64, index int, value float64) (newPtr []float64)

// ----------------------------------------------------------------------------

// 获取指定数据最先出现的位置的索引
// get the first index of target itme.
// 指定要素が最初に現れる場所のインデックスを取得
elarr.Index(v []interface{}, item interface{}) int
elarr.IndexInter(v []interface{}, item interface{}) int
elarr.IndexStr(v []string, item string) int
elarr.IndexInt(v []int, item int) int
elarr.IndexInt8(v []int8, item int8) int
elarr.IndexInt16(v []int16, item int16) int
elarr.IndexInt32(v []int32, item int32) int
elarr.IndexInt64(v []int64, item int64) int
elarr.IndexUint(v []uint, item uint) int
elarr.IndexUint8(v []uint8, item uint8) int
elarr.IndexUint16(v []uint16, item uint16) int
elarr.IndexUint32(v []uint32, item uint32) int
elarr.IndexUint64(v []uint64, item uint64) int
elarr.IndexRune(v []rune, item rune) int
elarr.IndexByte(v []byte, item byte) int
elarr.IndexFloat32(v []float32, item float32) int
elarr.IndexFloat64(v []float64, item float64) int

// ----------------------------------------------------------------------------

// 获取指定数据最后出现的位置的索引
// get the last index of target itme.
// 指定要素が最後に現れる場所のインデックスを取得
elarr.LIndex(v []interface{}, item interface{}) int
elarr.LIndexInter(v []interface{}, item interface{}) int
elarr.LIndexStr(v []string, item string) int
elarr.LIndexInt(v []int, item int) int
elarr.LIndexInt8(v []int8, item int8) int
elarr.LIndexInt16(v []int16, item int16) int
elarr.LIndexInt32(v []int32, item int32) int
elarr.LIndexInt64(v []int64, item int64) int
elarr.LIndexUint(v []uint, item uint) int
elarr.LIndexUint8(v []uint8, item uint8) int
elarr.LIndexUint16(v []uint16, item uint16) int
elarr.LIndexUint32(v []uint32, item uint32) int
elarr.LIndexUint64(v []uint64, item uint64) int
elarr.LIndexRune(v []rune, item rune) int
elarr.LIndexByte(v []byte, item byte) int
elarr.LIndexFloat32(v []float32, item float32) int
elarr.LIndexFloat64(v []float64, item float64) int

// ----------------------------------------------------------------------------

// 从切片删除第一个和指定内容一致的数据
// remove the first element that is same as the given item
// スライスから最初に与えられたデータと等価の要素を削除する
elarr.Remove(v []interface{}, item interface{}) (newPtr []interface{})
elarr.RemoveInter(v []interface{}, item interface{}) (newPtr []interface{})
elarr.RemoveStr(v []string, item string) (newPtr []string)
elarr.RemoveInt(v []int, item int) (newPtr []int)
elarr.RemoveInt8(v []int8, item int8) (newPtr []int8)
elarr.RemoveInt16(v []int16, item int16) (newPtr []int16)
elarr.RemoveInt32(v []int32, item int32) (newPtr []int32)
elarr.RemoveInt64(v []int64, item int64) (newPtr []int64)
elarr.RemoveUint(v []uint, item uint) (newPtr []uint)
elarr.RemoveUint8(v []uint8, item uint8) (newPtr []uint8)
elarr.RemoveUint16(v []uint16, item uint16) (newPtr []uint16)
elarr.RemoveUint32(v []uint32, item uint32) (newPtr []uint32)
elarr.RemoveUint64(v []uint64, item uint64) (newPtr []uint64)
elarr.RemoveRune(v []rune, item rune) (newPtr []rune)
elarr.RemoveRune(v []rune, item rune) (newPtr []rune)
elarr.RemoveFloat32(v []float32, item float32) (newPtr []float32)
elarr.RemoveFloat64(v []float64, item float64) (newPtr []float64)

// ----------------------------------------------------------------------------

// 其他函数
// other functions
// その他
elarr.Sum(v []interface{}) float64
elarr.SumInt(v []int) int
elarr.Average(v []interface{}) float64
elarr.Insert(v []interface{}, index int, value interface{}) (newPtr []interface{})
elarr.Join(v []interface{}, sep string) string
elarr.IsEmpty(v interface{}) bool
```
