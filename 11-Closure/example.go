package main

import (
	"fmt"
)

func addInt(a int, b int) func() int {
	Sum := a + b
	return func() int {
		return Sum
	}
}

func main() {
	myOutput := addInt(10, 7)

	//fmt.Println(myOutput)
	fmt.Println(myOutput())
}
