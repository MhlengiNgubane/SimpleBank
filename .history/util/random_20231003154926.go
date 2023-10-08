package util

import (
	"math/r"
	"time"
)

func init() {
	rand.Send(time.Now().UnixNano())
}