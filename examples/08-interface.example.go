package main

import(
	"fmt"
)

func emptyInterfaceExamples(){
	var emptyInterface interface{}
	emptyInterface = 100
	emptyInterface = "Hello, interface"

	fmt.Println(emptyInterface)

	emptyInterface = 1.343936
	fmt.Println(emptyInterface)
}

func switchingTypesExamples(){
	var aInterface interface{}
	aInterface = "Nothing"

	switch aInterface.(type) {
	case int:
		fmt.Println("This is an integer.")
	case string:
		fmt.Println("This is a string.")
	default:
		fmt.Println("This is an unknown type!!")
	}

}

/// Common interface example
type User struct {
	Name string
	Email string
	Username string
}

type Admin struct {
	User
	isAdmin bool
}

type UserInterface interface {
	GetAdminStatus() string
}

func (u User) GetAdminStatus() string{
	return fmt.Sprintf("The user is not a Admin!")
}

func (a Admin) GetAdminStatus() string{
	if a.isAdmin {
		return fmt.Sprintf("The user is an Admin!")
	}
	return fmt.Sprintf("The user is not a Admin!")
}
/// Common interface implementation ends

func main(){
	emptyInterfaceExamples()

	switchingTypesExamples()

	user := User{
		Name:    "John Doe",
		Email:    "john.doe@email.com",
		Username: "johnd",
	}

	admin := Admin{
		User:    User{
			Name:    "Chris Doe",
			Email:    "chris.doe@email.com",
			Username: "chrisd",
		},
		isAdmin: true,
	}

	nonAdmin := Admin{
		User:    User{
			Name:    "Chris Doe",
			Email:    "chris.doe@email.com",
			Username: "chrisd",
		},
		isAdmin: false,
	}

	fmt.Println(user.GetAdminStatus())
	fmt.Println(admin.GetAdminStatus())
	fmt.Println(nonAdmin.GetAdminStatus())

}
