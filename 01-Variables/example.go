package main

import (
	"fmt"
	"strings"
)

func main() {
	// Defining an empty string
	var myEmptyString string

	// Defining a string with initial value
	var myString = "Whatever"

	// Another way of defining a string with initial value
	str := "Third string"

	// Initializing a string array of size 5
	anotherString := make([]string, 5)

	fmt.Println(myEmptyString)

	fmt.Println(myString)

	fmt.Println(str)

	fmt.Println(anotherString)

	fmt.Println(len(anotherString))

	splitParamValue := strings.Split(myString, ",")
	fmt.Println("Split values: ", splitParamValue)
	fmt.Println("Split values length: ", len(splitParamValue))
}
