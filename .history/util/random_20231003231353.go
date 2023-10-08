package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInt 