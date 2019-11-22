package main

import(
	"fmt"
	"sync"
)

func printMe(){
	fmt.Println("Hello I am running from a Goroutine!")
}

func printMeWithItem(item string){
	fmt.Printf("Hello I am '%s' running from a Goroutine! \n", item)
}

func main(){
	var wg sync.WaitGroup

	// First example of go routine
	wg.Add(1)
	go func(){
		printMe()
		wg.Done()
	}()
	wg.Wait()

	// Second example for go routine in loops
	myList := []string{"Nothing", "Everything", "Whatever" }
	for _, item := range myList{
		wg.Add(1)
		go func(item string){
			printMeWithItem(item)
			wg.Done()
		}(item)
	}

	wg.Wait()


	/*
	Using an empty select{} can be used to keep function alive indefinitely
	*/
	//	select{}
}
