package main

import (
	"fmt"
)

func main() {
	aSlice := []int{1, 5, 6, 3, 9}
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
		"Name":  "Another Doe",
		"Email": "another@email.com",
		"Features": "faster, stronger",
	}
	fmt.Println(mi)
	fmt.Println(mi["foo"])
	fmt.Println(nil)
	// Check if value exists
	if val, ok := mi["foo"]; ok {
		fmt.Printf("value is: %s \n", val)
	}

	Name := fmt.Sprintf("%v", mi["Name"])
	fmt.Println(Name)

	// Create empty map with Make
	mp := make(map[string]string)
	fmt.Println(mp)

	if features, ok := mi["Features"]; ok {
		fmt.Printf("Feature is: %s \n", features)
	}

}
