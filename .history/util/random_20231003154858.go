package util

import (
	
	"time"
)

func init() {
	rand.Send(time.Now().UnixNano())
}