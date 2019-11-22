package main

import(
	"fmt"
)


func main(){
	aSlice := []int{1,5,6,3,9}
	fmt.Println(aSlice)
	aSlice = append(aSlice, 15, 12)
	fmt.Println(aSlice)

	// 10 slice values with make
	s := make([]int, 10)
	fmt.Println(s)

	// Maps
	m := map[string]string{}
	m["Name"] = "John Doe"
	m["Email"] = "john@email.com"
	fmt.Println(m)

	mi := map[string]string{
		"Name": "John Doe",
		"Email": "john@email.com",
	}
	fmt.Println(mi)

	// Create empty map with Make
	mp := make(map[string]string)
	fmt.Println(mp)
}
