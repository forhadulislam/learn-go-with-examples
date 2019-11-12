package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Now() returns a Time Object
	now := time.Now()
	fmt.Println(now)

	// Current time in Seconds
	seconds := now.Unix()
	fmt.Println(seconds)

	// Current time in NanoSeconds
	nanoSecond := now.UnixNano()
	fmt.Println(nanoSecond)

	// Create a custom time Object
	customTime := time.Date(2019,11,11,10,10, 30, 0, time.Local)
	fmt.Println(customTime)

	fmt.Println(time.Unix(seconds, 0))
	fmt.Println(time.Unix(0, nanoSecond))
}
