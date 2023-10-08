package util

import (
	"math/rand"
	"time"
)

func init() {r := rand.New(rand.NewSource(VALUE))

	rand.Seed(time.Now().UnixNano())
}