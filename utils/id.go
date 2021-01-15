package utils

import (
	"strconv"
	"time"
)


var (
	DefMinInt = 1000
	DefMaxInt = 9999
)

// MicroTimeID generate.
// return like: 16074145697981929446(len: 20)
func MicroTimeID() string {
	ms := time.Now().UnixNano() / 1000
	ri := RandomInt(DefMinInt, DefMaxInt)
	return strconv.FormatInt(ms, 10) + strconv.FormatInt(int64(ri), 10)
}
