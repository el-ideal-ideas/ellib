package elconv

import "github.com/el-ideal-ideas/ellib/elstr"

// convert value to string.
func AsStr(v interface{}) string {
	s, err := elstr.ToString(v)
	if err != nil {
		return ""
	} else {
		return s
	}
}
