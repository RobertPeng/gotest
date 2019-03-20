package main

import (
	"fmt"
	"time"
)

func init() {
	Timer()
}

func Timer() {
	for {
		select {
		case <-time.After(10 * time.Microsecond):
			fmt.Println("after timer")
		}
	}
}
