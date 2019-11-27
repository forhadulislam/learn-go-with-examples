package main

import (
	"fmt"
	"time"
)

func backgroundTicker() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Tock")
	}
}

func main() {

	//go backgroundTicker()
	timeP := time.Now()
	ticker := time.NewTicker(100 * time.Millisecond)

	done := make(chan bool, 1)

	myTicker := time.NewTicker( 100 * time.Millisecond)
	fmt.Println(myTicker)

	// Ticker example with Time difference
	Loop:
	for {
		select {
		case <- done:
			fmt.Println("2 seconds passed")
			break Loop
		case t := <-ticker.C:
			timeN := time.Now()
			diff := timeN.Sub(timeP).Seconds()

			if diff >= 2 {
				myTicker.Stop()
				fmt.Println("2 seconds passed channel is true")
				done <- true
			}
			fmt.Println("myTicker at ", t)
			fmt.Println("Diff ", diff)


		}
	}

	// Ticker example with time.After
	anotherTicker := time.NewTicker( 100 * time.Millisecond)
	fmt.Println(anotherTicker)
	timeOut := time.After(2 * time.Second)
	for {
		select {
		case <- timeOut:
			fmt.Println("2 seconds passed")
			return
		case t := <-ticker.C:
			fmt.Println("anotherTicker at ", t)


		}
	}

}
