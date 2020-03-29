package main

import (
	"fmt"
	"sync"
)

func main() {
	myt := sync.Mutex{}
	data := make(map[string]string)

	myt.Lock()
	fmt.Println("Print anything", data)
	myt.Unlock()
}
