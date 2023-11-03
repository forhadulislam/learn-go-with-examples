package main

import (
	"fmt"
)

//go:generate  mockgen -destination=mock_api.go -package=main -source=mock.go API
type API interface {
	CreateItem(data map[string]any) (id string, err error)
	GetItem(id string) (data map[string]any, err error)
	DeleteItem(id string) error
}

func main() {

	fmt.Println("Hello, playground")
}
