package main

import (
	"fmt"
	"math/rand"
	"net"
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


func getFreePort() int {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
	return listener.Addr().(*net.TCPAddr).Port
}

func RandomNumber(length int) string {
	return string(rand.Intn(100))
}

func main() {
	randomString := RandomString(50)
	fmt.Println("Generated random string : ", randomString)
	fmt.Println("Generated random number : ", RandomNumber(16))

	freePort := getFreePort()
	fmt.Println("My free port is ", freePort)
}
