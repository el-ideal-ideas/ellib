package utils

import (
	"math/rand"
	"time"
)


// RandomInt return a random int in [min,max)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
