## elrand
- 关于随机的工具包
- utilities for random
- 乱数関連のユーティリティ
```go
import "github.com/el-ideal-ideas/ellib/elrand"

// 获取指定范围的随机数(左包含，右不包含)
// random int in [min,max)
// 指定範囲の乱数を生成(min以上, maxより小さい)
elrand.RandomInt(min, max int) int

// 获取指定长度的随机的字符串
// get a random string.
// 指定の長さのランダムな文字列を生成
elrand.RandomString(length int) string

// 打乱数组内容
// shuffle the slice.
// スライスの内容をランダムにシャッフルする
elrand.Shuffle(arr []interface{})

// 从数组内随机选取指定长度的数据
// choice some random item from data.
// 指定の数のデータをランダムに選択する
elrand.Choice(data interface{}, count int) []interface{}
elrand.ChoiceInter(data []interface{}, count int) []interface{}
elrand.ChoiceStr(data []string, count int) []string
elrand.ChoiceInt(data []int, count int) []int
elrand.ChoiceInt8(data []int8, count int) []int8
elrand.ChoiceInt16(data []int16, count int) []int16
elrand.ChoiceInt32(data []int32, count int) []int32
elrand.ChoiceInt64(data []int64, count int) []int64
elrand.ChoiceUint(data []uint, count int) []uint
elrand.ChoiceUint8(data []uint8, count int) []uint8
elrand.ChoiceUint16(data []uint16, count int) []uint16
elrand.ChoiceUint32(data []uint32, count int) []uint32
elrand.ChoiceUint64(data []uint64, count int) []uint64
elrand.ChoiceRune(data []rune, count int) []rune
elrand.ChoiceByte(data []byte, count int) []byte
elrand.ChoiceFloat32(data []float32, count int) []float32
elrand.ChoiceFloat64(data []float64, count int) []float64
elrand.ChoiceBool(data []bool, count int) []bool
```

- elrand包含官方rand包的全部内容。
- elrand contains all functions in official rand package.
- elrandは公式のrandパッケージの全ての関数を含んでいます。

```go
type Rand = rand.Rand
type Source = rand.Source
type Source64 = rand.Source64
type Zipf = rand.Zipf

var Int = rand.Int
var Intn = rand.Intn
var Float64 = rand.Float64
var Float32 = rand.Float32
var Perm = rand.Perm
var Read = rand.Read
var ExpFloat64 = rand.ExpFloat64
var Int31a = rand.Int31
var Int63 = rand.Int63
var Int31n = rand.Int31n
var Int63n = rand.Int63n
var New = rand.New
var NewSource = rand.NewSource
var NewZipf = rand.NewZipf
var NormFloat64 = rand.NormFloat64
var Seeda = rand.Seed
var RShuffle = rand.Shuffle
var Uint32 = rand.Uint32
var Uint64 = rand.Uint64
```