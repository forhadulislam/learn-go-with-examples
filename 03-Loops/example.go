package main

import (
	"fmt"
)

func main() {
	// Basic foor loop example
	for i := 0; i <= 10; i++ {
		fmt.Printf("Hello, my index is '%d' running inside a loop! \n", i)
	}

	myList := []string{"Sam", "Ahmad", "Shihab", "Sami"}

	var err error
	for data := range myList {
		if myList[data] == "Shihab" {
			err = fmt.Errorf("blocked user")
		}
		fmt.Printf("What name is it '%s'!\n", myList[data])
	}
	if err != nil {
		fmt.Printf("You have got an error '%s'!\n", err)
	}

	// For loop example with traditional way
	for i := 0; i < len(myList); i++ {
		fmt.Printf("Hello I am '%s' running inside a for loop! \n", myList[i])
	}

	// For loop example with Range
	for _, item := range myList {
		fmt.Printf("Hello I am '%s' running inside a for loop with range! \n", item)
	}

	// Infinity for loop with a break
	var count = 9
	for {
		if count == 0 {
			break
		} else {
			fmt.Printf("Hello my number is '%d', I am still counting! \n", count)
		}
		count--
	}

	myListDuplicate := []string{"ABC", "DEF"}
	fmt.Print(myListDuplicate)

	for ind := range myListDuplicate {
		fmt.Println(myListDuplicate[ind])
	}
}
