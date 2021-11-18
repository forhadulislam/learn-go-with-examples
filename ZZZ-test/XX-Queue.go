package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func main() {
	subArray := sortArray([]int{2, 7, 1, 3, 4, 5, 9, 14, 19})
	sort.Ints(subArray) // built-in go method to sort array of integers
	fmt.Println(subArray)
}

func sortSubArray(arr []int) []int {
	sort.Ints(arr) // built-in go method to sort array of integers
	fmt.Println(arr)
	return arr
}

func sortArray(x []int) (sortedSubArray []int) {
	length := len(x)
	increment := 4
	if length == 0 {
		return
	}

	for i:=0; i < (length); i += increment  {
		wg.Add(1)
		subArr := []int{}
		if i + increment > length - 1 {
			subArr = x[i:]
		}else{
			subArr = x[i:i+increment]
		}
		sortedSubArray = append(sortedSubArray, subArr...)
		fmt.Printf("sub-array: %v \n", subArr)
		go func() {
			sortSubArray(subArr)
			wg.Done()
		}()
	}
	wg.Wait()

	return sortedSubArray
}
