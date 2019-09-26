package main

import(
	"fmt"
)

func main(){

	ch := make(chan int)
	go func() {
		ch <- 19
	}()

	go func() {
		ch <- 19
		val := <- ch
		fmt.Println(val)
	}()

}
