package util

import (
	"math/rand"
	"time"
)

func init() {
	NewRand(NewSource(seed(time.Now().UnixNano()))
}