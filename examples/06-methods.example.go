package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const characterOptions = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byte := make([]byte, length)
	for i := range byte {
		byte[i] = characterOptions[rand.Intn(len(characterOptions))]
	}
	return string(byte)
}

func main() {
	randomString := RandomString(50)
	fmt.Println("Generated random string : ", randomString)
}
