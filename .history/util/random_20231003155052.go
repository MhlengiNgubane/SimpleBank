package util

import (
	"math/rand"
	"time"
)

func init() {
	rand(time.Now().UnixNano())
}