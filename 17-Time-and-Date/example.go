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
	customTime := time.Date(2019, 11, 11, 10, 10, 30, 0, time.Local)
	fmt.Println(customTime)
	fmt.Println(customTime.UTC().Format(time.RFC3339))

	fmt.Println(time.Unix(seconds, 0))
	fmt.Println(time.Unix(0, nanoSecond))

	cTime, _ := time.Parse("2019-11-13T13:04:08Z", "2019-11-13T13:04:08Z")
	fmt.Println(cTime)

	// Calculate time difference
	previousTime := time.Date(2019, 11, 11, 10, 10, 30, 0, time.Local)
	timeDifference := now.Sub(previousTime).Nanoseconds()
	println("Time difference is: ", timeDifference)

	// Location
	tz := time.Now().Location().String()
	fmt.Println(tz)

	// Timezone
	loc, _ := time.LoadLocation("UTC")
	now = time.Now().In(loc)
	fmt.Println("ZONE : ", loc, " Time : ", now) // UTC

	t := time.Now()
	zone, offset := t.Zone()
	fmt.Println(zone, offset)
}
