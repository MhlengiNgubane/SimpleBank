package util

import (
	"math/rand"
	"time"
)

func init() {
	NewSource(seed(time.Now().UnixNano())
}