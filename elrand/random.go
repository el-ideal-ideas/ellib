package elrand

import (
	"github.com/el-ideal-ideas/ellib/elstr"
	"math/rand"
	"time"
)

const (
	rsLetters       = elstr.AsciiLetters + elstr.Digits
	rsLetterIdxBits = 6
	rsLetterIdxMask = 1<<rsLetterIdxBits - 1
	rsLetterIdxMax  = 63 / rsLetterIdxBits
)

// RandomInt returns a random int in [min,max)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandomString returns a random string.
func RandomString(length int) string {
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	cache, remain := rand.Int63(), rsLetterIdxMax
	for i := length - 1; i >= 0; {
		if remain == 0 {
			rand.Seed(time.Now().UnixNano())
			cache, remain = rand.Int63(), rsLetterIdxMax
		}
		idx := int(cache & rsLetterIdxMask)
		if idx < len(rsLetters) {
			b[i] = rsLetters[idx]
			i -= 1
		}
		cache >>= rsLetterIdxBits
		remain -= 1
	}
	return string(b)
}

// Shuffle array or slice.
func Shuffle(arr []interface{}) {
	rand.Seed(time.Now().UnixNano())
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
