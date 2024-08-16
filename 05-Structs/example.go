package main

import (
	"fmt"
)

// Maps are good but they have limitations.
// They don’t define an API since there’s no way to constrain a map to only allow certain keys.
// Also, all the values in a map must be of the same type.

// Example 1: Basics of Struct

type Vehicle struct {
	Brand    string
	Model    string
	Year     int
	Warranty int
}

type People struct {
	Username string
	Email    string
	age      int
}

type Client struct {
	Company string
	People
	V Vehicle
}

func (v *Vehicle) printMyVehicle() {
	fmt.Println("printMyVehicle() starts")
	fmt.Println(v)
	fmt.Println("printMyVehicle() ends")
}


type MyString string

func (r MyString) String() string {
	return string(r)
}

// Main function
func main() {

	a := MyString("Hello")
	fmt.Println("Value for a:", a.String())

	// Part of Example 1
	v := Vehicle{
		Brand:    "Hyundai",
		Model:    "i30",
		Year:     2019,
		Warranty: 10,
	} // This is a map literal style of struct definition. Struct literal is without fields (i.e Brand, Model).
	v.printMyVehicle()

	annonymusStruct := struct {
		Name  string
		Email string
		Phone string
	}{
		"John",
		"john.doe@email.com",
		"+187778AAA",
	}

	fmt.Println(annonymusStruct)

	c := Client{
		Company: "Apple",
		People: People{
			"john",
			"john@doe.com",
			29,
		},
		V: Vehicle{},
	}

	fmt.Println(c)
}
