package main

import (
	"fmt"
	"reflect"
)

type Library struct {
	books string
	description string
}

func convert(val interface{}){
	m := reflect.ValueOf(val)
	fmt.Print(m)
}

func main() {

	lib := Library{
		books:       "Nothing to lose",
		description: "You have nothing to lose",
	}

	convert(lib)
}
