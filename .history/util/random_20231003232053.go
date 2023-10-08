package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string on length n
func RandomString(n int) string