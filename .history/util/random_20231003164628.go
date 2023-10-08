package util

import (
	"math/rand"
	"time"
)

func (r *Rand) init() {
	r := rand.New(rand.NewSource(seed))
}