package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.S(time.Now().UnixNano())
}