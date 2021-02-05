## elstr
- 用于处理字符串的工具包
- utilities for string
- 文字列処理用のユーティリティ

#### strings
- elstr包含官方strings包的全部函数 (为了性能优化一部分函数用了其他的实现方式)
- elstr contains all functions in official strings package (some functions use different source code for performance)
- elstrパブリックは公式のstringsパッケージの全部の関数を含んでいます (一部の関数は性能向上のため、異なる実装になっています)
- [strings](https://github.com/el-ideal-ideas/ellib/blob/master/doc/strings.md)

#### xstrings(https://github.com/huandu/xstrings)
- elstr包含xstrings包的全部函数
- elstr contains all functions in xstrings package
- elstrはxstringsパッケージの全部の関数を含んでいます
- [xstrings](https://github.com/el-ideal-ideas/ellib/blob/master/doc/xstrings.md)

##### constants in python
```go
// -------- these constants are from python's string module --------

// Whitespace -- a string containing all ASCII whitespace
// AsciiLowercase -- a string containing all ASCII lowercase letters
// AsciiUppercase -- a string containing all ASCII uppercase letters
// AsciiLetters -- a string containing all ASCII letters
// Digits -- a string containing all ASCII decimal digits
// HexDigits -- a string containing all ASCII hexadecimal digits
// OctDigits -- a string containing all ASCII octal digits
// PUNCTUATION -- a string containing all ASCII punctuation characters
// Printable -- a string containing all ASCII characters considered printable
const Whitespace = " \t\n\r\v\f"
const AsciiLowercase = "abcdefghijklmnopqrstuvwxyz"
const AsciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const AsciiLetters = AsciiLowercase + AsciiUppercase
const Digits = "0123456789"
const HexDigits = Digits + "abcdef" + "ABCDEF"
const OctDigits = "01234567"
const PUNCTUATION = `!"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`"
const Printable = Digits + AsciiLetters + PUNCTUATION + Whitespace
```

##### functions in python
```go
import "github.com/el-ideal-ideas/ellib/elstr"

// 检测字符串是否只包含英数字
// return true if the string is an alpha-numeric string
// 文字列が英数字のみであれば真を返す
elstr.IsAlNum(s string) bool

// 检测字符串是否只包含英字
// return true if the string is an alphabetic string
// 文字列が英字のみであれば真を返す
elstr.IsAlpha(s string) bool

// 检测字符串是否只包含数字
// return true if the string is a digit string
// 文字列が数字のみであれば真を返す
elstr.IsDigit(s string) bool

// 检测字符串是否只包含ascii编码内的字符
// return true if all characters in the string are ASCII
// 文字列がアスキー文字のみであれば真を返す
elstr.IsASCII(s string) bool

// 检测字符串内的英文字符是不是都是大写
// return true if the string is an uppercase string
// 文字列内の英字が全部大文字であれば真を返す
elstr.IsUpper(s string) bool

// 检测字符串内的英文字符是不是都是小写
// return true if the string is an lowercase string
// 文字列内の英字が全部小文字であれば真を返す
elstr.IsLower(s string) bool

// 去除字符串的头部和末尾的空格
// return a slice of the string s, with all leading and trailing white space removed
// 先頭と末尾の空白を削除
elstr.Strip(s string) string
```

##### other functions
```go
import "github.com/el-ideal-ideas/ellib/elstr"

// 置换字符串但是返回[]byte
// replace strings but return []byte
// 文字列置換関数、[]byteを返す
elstr.ReplaceToBytes(s, old, new string, n int) []byte

// 循环字符串但是返回[]byte
// repeat the string but return []byte
// 文字列を繰り返し、[]byteを返す
elstr.RepeatToBytes(s string, count int) []byte

// 拼接字符串但是返回[]byte
// join strings but return []byte
// 文字列連結する、[]byteを返す
elstr.JoinToBytes(a []string, sep string) []byte

// 字符串转换成[]byte (比 []byte("some string") 要快一点)
// convert string to []byte (faster than []byte("some string"))
// 文字列を[]byteに変換 ([]byte("some string")より速い)
elstr.ToBytes(s string) []byte

// 根据自然语言排序字符串
// sorts strings according to their natural sort order
// 自然順ソート
elstr.SortStringsNaturally(s []string)

// 检测字符串的相似度
// similarity calc for two string
// 文字列の類似度を計算
elstr.Similarity(s, t string, rate float32) (float32, bool)

// fmt.Sprintf的别名
// alias of fmt.Sprintf
// fmt.Sprintfの別名
elstr.Format(format string, a ...interface{}) string

// 字符串反转
// reverse string
// 文字列反転
elstr.Reverse(s string) string

// 字符串反转 (只对应ascii的字符串)
// reverse string (only supports ascii strings)
// 文字列反転 (アスキー文字のみ対応)
elstr.ReverseASCII(s string) string

// 字符串置换 (比strings.Replace快一些)
// replace string (faster than strings.Replace)
// 文字列置き換え (strings.Replaceより速い)
elstr.Replace(s, old, new string, n int) string

// 字符串循环 (比strings.Repeat快一些)
// repeat string (faster than strings.Repeat)
// 文字列繰り返し (strings.Repeatより速い)
elstr.Repeat(s string, count int) string

// 字符串连接 （比strings.Join快一些）
// Join strings (faster than strings.Join)
// 文字列連結　(strings.Joinより速い)
elstr.Join(a []string, sep string) string

// 把数值当做字符串拼接
// Join numbers as strings
// 数値を文字列として連結
elstr.JoinInts(ints []int, sep string) (res string)

// 复制字符串
// copy string
// 文字列コピー
elstr.Copy(src string) string

// 剪切字符串（支持rune）
// get a part of string (supports rune)
// 文字列切り取り (runeをサポート)
elstr.SubString(s string, start, length int) string

// 分隔字符串（支持多个切割符号）
// split strings (supports multiple separators)
// 文字列分割 (複数の区切り文字をサポート)
elstr.SplitMulti(s string, sep ...string) (ss []string)

// 字符串置换
// replace string
// 文字列置き換え
elstr.Replaces(str string, pairs map[string]string) string

// 把16进制的数字，转换成数值
// convert hex string to int
// 16進数の文字列を数値に変換
elstr.Hex2Dec(str string) (int64, error)

// 检测字符串是否是合法的邮箱格式
// if `s` is a valid format as email address, return true
// 文字列が正しいメールアドレスの形式であれば真を返す
elstr.IsEmailFormat(s string) bool

// 如果所有字符都是中文，返回true
// if all characters in `s` is a Chinese character, return true
// 全ての文字が中国語であれば真を返す
elstr.IsAllChineseChar(s string) bool

// 如果是回文就返回true
// if `s` is palindrome, return true
// 回文であれば真を返す
elstr.IsPalindrome(s string) bool

// 尝试转换成字符串
// try convert to string
// 文字列への変換を試みる
elstr.ToString(val interface{}) (str string, err error)

// 强制转换成字符串
// convert to string
// 文字列に強制変換
elstr.MustString(v interface{}) string

// 返回地n个rune的index
// get the index of `n`th rune character.
// n番目のruneのindexを取得
elstr.RuneIndex(p []byte, n int) (int, bool)
elstr.RuneIndexInString(s string, n int) (int, bool)

// 切割字符串（支持rune）
// get a part from string(supports rune)
// 文字列切り取り (runeをサポート)
elstr.RuneSub(p []byte, start, length int) []byte
elstr.RuneSubString(s string, start, length int) string
```

##### Unsafe
- 这些函数虽然速度特别快，但是具有一定危险性，使用前请阅读源码
- Following functions is faster but dangerous, Please read the source code before use it
- 下記の関数は処理速度が高速であるが、特定の場面でエラーを起こす危険性があり、使用する前にソースコードを確認してください。
```go
import "github.com/el-ideal-ideas/ellib/elstr"

elstr.UnsafeReplaceToBytes(s, old, new string, n int) []byte
elstr.UnsafeToBytes(s string) []byte
elstr.UnsafeToString(s []byte) string
```

##### constants
```go
// japanese Hiragana list
// あ -- ん [:46]
// ぁ -- っ [46:55]
// が -- ぽ [55:]
var Hiragana = []rune{
	'あ', 'い', 'う', 'え', 'お',
	'か', 'き', 'く', 'け', 'こ',
	'さ', 'し', 'す', 'せ', 'そ',
	'た', 'ち', 'つ', 'て', 'と',
	'な', 'に', 'ぬ', 'ね', 'の',
	'は', 'ひ', 'ふ', 'へ', 'ほ',
	'ま', 'み', 'む', 'め', 'も',
	'や', 'ゆ', 'よ',
	'ら', 'り', 'る', 'れ', 'ろ',
	'わ', 'を',
	'ん',
	'ぁ', 'ぃ', 'ぅ', 'ぇ', 'ぉ',
	'ゃ', 'ゅ', 'ょ',
	'っ',
	'が', 'ぎ', 'ぐ', 'げ', 'ご',
	'ざ', 'じ', 'ず', 'ぜ', 'ぞ',
	'だ', 'ぢ', 'づ', 'で', 'ど',
	'ば', 'び', 'ぶ', 'べ', 'ぼ',
	'ぱ', 'ぴ', 'ぷ', 'ぺ', 'ぽ',
}

// japanese Katakana list
// ア -- ン [:46]
// ァ -- ッ [46:55]
// ガ -- ポ [55:]
var Katakana = []rune{
	'ア', 'イ', 'ウ', 'エ', 'オ',
	'カ', 'キ', 'ク', 'ケ', 'コ',
	'サ', 'シ', 'ス', 'セ', 'ソ',
	'タ', 'チ', 'ツ', 'テ', 'ト',
	'ナ', 'ニ', 'ヌ', 'ネ', 'ノ',
	'ハ', 'ヒ', 'フ', 'ヘ', 'ホ',
	'マ', 'ミ', 'ム', 'メ', 'モ',
	'ヤ', 'ユ', 'ヨ',
	'ラ', 'リ', 'ル', 'レ', 'ロ',
	'ワ', 'ヲ',
	'ン',
	'ァ', 'ィ', 'ゥ', 'ェ', 'ォ',
	'ャ', 'ュ', 'ョ',
	'ッ',
	'ガ', 'ギ', 'グ', 'ゲ', 'ゴ',
	'ザ', 'ジ', 'ズ', 'ゼ', 'ゾ',
	'ダ', 'ヂ', 'ヅ', 'デ', 'ド',
	'バ', 'ビ', 'ブ', 'ベ', 'ボ',
	'パ', 'ピ', 'プ', 'ペ', 'ポ',
}

// alphabet uppercase letter
var AlphabetUpper = []rune{
	'A', 'B', 'C', 'D', 'E', 'F', 'G',
	'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U',
	'V', 'W', 'X', 'Y', 'Z',
}

// alphabet lowercase letter
var AlphabetLower = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u',
	'v', 'w', 'x', 'y', 'z',
}

// number list
var Number = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// chinese number list
var ChineseDigitsTraditional = []rune{'零', '壹', '贰', '叁', '肆', '伍', '陆', '柒', '捌', '玖'}
var ChineseDigitsSimple = []rune{'〇', '一', '二', '三', '四', '五', '六', '七', '八', '九'}

// japanese number list
var JapaneseDigitsHiragana = []string{"いち", "に", "さん", "し", "ご", "ろく", "しち", "はち", "きゅう", "じゅう"}
var JapaneseDigitsKatakana = []string{"イチ", "ニ", "サン", "シ", "ゴ", "ロク", "シチ", "ハチ", "キュウ", "ジュウ"}
```