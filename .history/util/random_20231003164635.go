package util

import (
	"math/rand"
	"time"
)

func (r *Rand) init() in {
	r := rand.New(rand.NewSource(seed))
}