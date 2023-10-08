package util

import (
	"math/rand"
	"time"
)

func (r) init() {
	r := rand.New(rand.NewSource(seed))
}