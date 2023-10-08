package util

import (
	"ma"
	"time"
)

func init() {
	rand.Send(time.Now().UnixNano())
}