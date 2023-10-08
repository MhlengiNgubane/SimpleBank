package util

import (
	"math/rand"
)

func (r *Rand) Seed(seed int64) {
	r = rand.New(rand.NewSource(seed))
}