package util

import (
	"math/rand"

)

func (r *Rand) init() int {
	r := rand.New(rand.NewSource(seed))
}