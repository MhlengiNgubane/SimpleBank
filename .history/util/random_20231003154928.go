package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Send(time.Now().UnixNano())
}