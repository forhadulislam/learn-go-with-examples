package main

import (
	"fmt"
	"reflect"
	"strings"
)

func updateMap(data map[string]interface{}, key string) map[string]interface{} {
	var dataClone = data
	//dataClont = copy(dataClone, data)
	dataClone[key] = "********"
	return dataClone
}

func copyMap(originalMap map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}

func filterMap(originalMap map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}

func main() {
	aSlice := []int{1, 5, 6, 3, 9}
	fmt.Println(aSlice)

	aSlice = append(aSlice, 15, 12)
	fmt.Println(aSlice)

	// If you have a sparse array (an array where most elements are set to their zero value),
	// you can specify only the indices with values in the array literal:
	var x = [12]int{1, 5: 4, 6, 9: 100, 15}
	fmt.Println(x)

	var a [2][3]int
	fmt.Println(a)
	fmt.Println(len(a))

	// 10 slice values with make
	s := make([]int, 10)
	fmt.Println(s)

	// Maps
	m := map[string]interface{}{}
	m["Name"] = "John Doe"
	m["Email"] = "john@email.com"
	m["Password"] = "@veryStrongP@$$W0RD"
	mCopy := copyMap(m)
	fmt.Println(m)
	fmt.Println(updateMap(m, "Password"))
	fmt.Println(m)

	//fmt.Println(updateMap(mCopy, "Password"))
	fmt.Println(mCopy)

	mi := map[string]string{
		"Name":     "Another Doe",
		"Email":    "another@email.com",
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

	// Complicated map
	var aBigMap = map[string]map[string]map[string]string{
		"user1": {
			"normal": {
				"check":    "true",
				"redirect": "true",
			},
			"secure": {
				"check":    "true",
				"redirect": "false",
			},
		},
		"user2": {
			"http": {
				"check":    "true",
				"redirect": "true",
			},
			"https": {
				"check":    "true",
				"redirect": "false",
			},
		},
		"user3": {
			"default": {
				"check":    "true",
				"redirect": "false",
			},
		},
	}

	fmt.Println(aBigMap["user3"]["default"])
	fmt.Println(aBigMap["user4"]["default"])
	fmt.Println(aBigMap["noOne"] == nil)
	fmt.Println(aBigMap["user3"]["default"] == nil)
	fmt.Println(aBigMap["user4"]["default"] == nil)

	stringArray := []string {"Hello","world","!"}
	justString := fmt.Sprintf("[%s]", strings.Join(stringArray, ", "))

	fmt.Println("value \t=", justString, "\ntype \t=", reflect.TypeOf(justString))
}
