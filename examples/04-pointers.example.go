package main

import(
	"fmt"
)

func main(){
	s := "first string"
	pointer := &s
	fmt.Println(s)
	fmt.Println(pointer)
	fmt.Println(*pointer)

	aValue := 20
	bValue := &aValue
	cV := aValue

	fmt.Println("// & is memory address and * is value reference")
	fmt.Println(aValue, bValue, *bValue, cV)

	aValue = 30
	fmt.Println(aValue, bValue, *bValue, cV)

	*bValue = 50
	fmt.Println(aValue, bValue, *bValue, cV)

	fmt.Println("// Creating a new integer")
	newInt := new(int)
	fmt.Println(newInt)
	fmt.Println(*newInt)

	fmt.Println("// *newInt = 55")
	*newInt = 55
	fmt.Println(newInt)
	fmt.Println(*newInt)

}
