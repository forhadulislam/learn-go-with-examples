package main

import(
	"fmt"
)

func main(){
	s := "first text"
	pointer := &s
	fmt.Println(s)
	fmt.Println(pointer)
	fmt.Println(*pointer)

	aValue := 20
	bValue := &aValue
	cV := aValue

	fmt.Println(aValue, bValue, *bValue, cV)

	aValue = 30
	fmt.Println(aValue, bValue, *bValue, cV)

	*bValue = 50
	fmt.Println(aValue, bValue, *bValue, cV)


	newInt := new(int)
	fmt.Println(newInt)
	fmt.Println(*newInt)

	*newInt = 55
	fmt.Println(newInt)
	fmt.Println(*newInt)

}
