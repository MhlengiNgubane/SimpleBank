package util

import (
	"math"
	"time"
)

func init() {
	rand.Send(time.Now().UnixNano())
}