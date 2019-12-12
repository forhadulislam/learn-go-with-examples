package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const characterOptions = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = characterOptions[rand.Intn(len(characterOptions))]
	}
	return string(bytes)
}

func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func setEnvVariables(envVars map[string]string) error {
	for k, v := range envVars {
		_, ok := os.LookupEnv(k)
		if !ok {
			err := os.Setenv(k, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RandomNumber(length int) string {
	return string(rand.Intn(100))
}

func main() {
	randomString := RandomString(50)
	fmt.Println("Generated random string : ", randomString)
	fmt.Println("Generated random number : ", RandomNumber(16))

	freePort, err := getFreePort()
	if err != nil {
		fmt.Println("Error getting a free port")
		os.Exit(1)
	}
	fmt.Println("My free port is ", freePort)
}
