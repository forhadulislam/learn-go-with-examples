package main

import (
	"errors"
	"fmt"
)

func fetchData(data string) (string, error) {
	if data != "invalid" {
		return "No error", nil
	}
	errorData := errors.New("invalid data received")
	return "", errorData
}

func main() {
	currentString, err := fetchData("Some data")
	fmt.Println("Current String :", currentString)
	fmt.Println("Error :", err)

	currentString, err = fetchData("invalid")
	fmt.Println("Current String :", currentString)
	fmt.Println("Error :", err)
}
