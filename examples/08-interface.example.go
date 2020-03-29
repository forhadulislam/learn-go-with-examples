package main

import (
	"fmt"
)

func emptyInterfaceExamples() {
	var emptyInterface interface{}
	emptyInterface = 100
	emptyInterface = "Hello, interface"

	fmt.Println(emptyInterface)

	emptyInterface = 1.343936
	fmt.Println(emptyInterface)
}

func switchingTypesExamples() {
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
type AdminInterface interface {
	AddAdmin() string
	GetAdmin() string
	RemoveAdmin() string
}

type Admin struct {
	Name         string
	Email        string
	Username     string
	isAdSupermin bool
}

type SuperUsers struct {
	Name        string
	Email       string
	Username    string
	isSuperUser bool
}

func (a Admin) AddAdmin() string {
	return fmt.Sprintf("added admin")
}

func (a Admin) GetAdmin() string {
	if a.isAdSupermin {
		return fmt.Sprintf("The user is super admin!")
	}
	return fmt.Sprintf("The user is a normal Admin!")
}

func (a Admin) RemoveAdmin() string {
	if a.isAdSupermin {
		return fmt.Sprintf("The user is super admin! cannot remove")
	}
	return fmt.Sprintf("admin is removed")
}

func (s SuperUsers) AddAdmin() string {
	return fmt.Sprintf("added user")
}

func (s SuperUsers) GetAdmin() string {
	if s.isSuperUser {
		return fmt.Sprintf("The user is super user!")
	}
	return fmt.Sprintf("The user is a normal user!")
}

func (s SuperUsers) RemoveAdmin() string {
	if s.isSuperUser {
		return fmt.Sprintf("The user is super user! cannot remove")
	}
	return fmt.Sprintf("user is removed")
}

func ExecuteCommonInterface(admin AdminInterface) {
	fmt.Println("Executing interface")
	fmt.Println(admin.AddAdmin())
	fmt.Println(admin.GetAdmin())
	fmt.Println(admin.RemoveAdmin())
	fmt.Println("execution ended")
}

/// Common interface implementation ends
func main() {
	emptyInterfaceExamples()
	switchingTypesExamples()

	admin := Admin{
		Name:     "John Doe",
		Email:    "john.doe@email.com",
		Username: "johnd",
	}

	sUser := SuperUsers{
		Name:        "John Doe",
		Email:       "john.doe@email.com",
		Username:    "johnd",
		isSuperUser: true,
	}

	ExecuteCommonInterface(admin)
	ExecuteCommonInterface(sUser)
}
