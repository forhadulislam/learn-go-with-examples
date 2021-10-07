package main

import (
	"time"
)

var name = "Some name"

func updateName(newName string) {
	name = newName
}


func main() {
	// Goroutine 1 is executed
	go updateName("Another Name")
	// Goroutine 2 is executed
	go updateName("Some other Name")
	// 1 second sleep is executed so all the goroutines can finish
	time.Sleep(1 * time.Second)
}

/*
	If we save this file as example.go and run this command go run -race 34.race-conditions/example.go,
	we will that there is a race condition as both of the goroutines can interfere 	while updating and
	reading the same variable `name`.
 */