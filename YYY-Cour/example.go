package main

import (
	"errors"
	"fmt"
	"time"
)

// User struct your colleague is implementing this
type User struct {
	ID string
	Name string
	Email string
}

func (u User) GetUserDetails(id string) (User, error) {
	if id != ""{
		return User{
			ID: id,
			Name: "Miikka",
		}, nil
	}
	return User{}, errors.New("no user id given")
}

func (u User) EditUserDetails(user User) error {
	if u.ID == ""{
		return errors.New("no user id given")
	}
	return nil
}

func (u User) DeleteUserDetails(id string) error {
	if u.ID != ""{
		return nil
	}
	return errors.New("no user id given")
}

// DBFuncUser is what you are writing as an interface
type DBFuncUser interface{
	GetUserDetails(id string) (User, error)
	EditUserDetails(user User) error
	DeleteUserDetails(id string) error
}

type AppContext struct {
	dbImpl DBFuncUser
}

func main(){
	u := User{}
	db := AppContext{
		dbImpl: u,
	}

	id := "19230213"
	user, err := db.dbImpl.GetUserDetails(id)
	if err != nil {
		fmt.Println( fmt.Errorf("%s", err) )
	}
	fmt.Println(user)

	str := "This is a string"
	myString(str)
	myStringPtr(&str)
	fmt.Printf("%s \n", str)
	fmt.Printf("%p \n", &str)
	//InterfaceExample()

	fmt.Println("New example ------ ")
	a := new(int)
	var abc int
	b := make([]int, 0)
	*a = 90
	fmt.Println("---a---")
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println("---b---")
	fmt.Println(b)
	fmt.Println("---abc---")
	fmt.Println(abc)
	fmt.Println(&abc)
	getValue(*a)
}

func getValue(in int)  {
	fmt.Println("---getValue---")
	str := "some value"
	fmt.Println("---str---")
	fmt.Println(str)
	fmt.Println(&str)
	fmt.Println("---in---")
	fmt.Println(in)
	fmt.Println(&in)
}

func myString(str string)  {
	fmt.Println(str)
	fmt.Printf("myString `str` pointer %p \n", &str)
	str = "Changed string"
}

func myStringPtr(str *string)  {
	fmt.Println(str)
	fmt.Printf("myStringPtr `str` pointer %p \n", str)
	*str = "Changed string"
}

type myStruct struct {
	myIntf myInterface
}

// Interface example
type ServerInterface struct {}

func (s ServerInterface) Start()  {
	fmt.Println("Server started")
	s.Run()
}

func (s ServerInterface) Stop()  {
	fmt.Println("Server stopped")
}

func (s ServerInterface) Run()  {
	fmt.Println("Server running")
}

type myInterface interface {
	Start()
	Stop()
}

func InterfaceExample() {
	server := new(ServerInterface)
	mystrct := myStruct{
		myIntf: server,
	}
	mystrct.myIntf.Start()
	time.Sleep(time.Second * 5)
	mystrct.myIntf.Stop()
}