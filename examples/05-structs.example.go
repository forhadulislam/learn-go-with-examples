package main

import(
	"fmt"
)

type Vehicle struct{
	Brand string
	Model string
	Year int
	Warranty int
}

type User struct{
	Username string
	Email string
	age int
}

type Client struct {
	Company string
	User
	V Vehicle
}

func(v *Vehicle) printMyVehicle(){
	fmt.Println("printMyVehicle() starts")
	fmt.Println(v)
	fmt.Println("printMyVehicle() ends")
}

func main() {
	v := Vehicle{
		Brand:    "Hyundai",
		Model:    "i30",
		Year:     2019,
		Warranty: 10,
	}
	v.printMyVehicle()

	annonymusStruct := struct {
		Name string
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
		User:    User{
			"john",
			"john@doe.com",
			29,
		},
		V:       Vehicle{
			
		},
	}

	fmt.Println(c)
	
}