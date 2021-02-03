## elenv
- 关于环境变量的包
- utilities for system environments
- 環境変数関連のユーティリティ
```go
import "github.com/el-ideal-ideas/ellib/elenv"

elenv.GetEnv(name string, def ...string) string
elenv.Setenv(key, value string) error
```