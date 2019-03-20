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
		case <-time.Tick(10 * time.Microsecond):
			fmt.Println("tick timer")
		}
	}
}
