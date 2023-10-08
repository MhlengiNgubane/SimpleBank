package util

import (
	"math/rand"
	"time"
)

func init() {
	NewRand(NewSource(seed))
	rand.Seed(time.Now().UnixNano())
}