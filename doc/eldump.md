## eldump
- 这个包可以把变量内容用容易理解格式打印到标准输出，方便开发和测试。
  这个包是从github.com/gookit/goutil复制过来的。
- This package can print data more clear and beautiful, For debug & development.
  This package was copy from github.com/gookit/goutil.
- 開発＆デバッグ用に変数の中身を見やすい形式で出力する。
  このパッケージはgithub.com/gookit/goutilからコピーされたものです。
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
