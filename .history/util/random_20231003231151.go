package util

import (
	"math/rand"
	"time"
)

func (r *Rand) init() int {
	r := rand.New(rand.NewSource(seed))
}