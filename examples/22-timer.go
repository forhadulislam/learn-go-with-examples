package main

import (
	"fmt"
	"time"
)

func main() {

	newTimer := time.NewTimer( time.Second )

	go func() {
		<-newTimer.C
		fmt.Println("Timer Goroutine expired")
	}()

	time.Sleep(time.Second * 2)
}
