package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
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

func getRandomIntInRange(min int, max int) int {
	if max <= 0 || min < 0 || min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
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

func isValidIP(ipAddress string) bool {
	parsedIP := net.ParseIP(ipAddress)
	isIPv4 := parsedIP != nil && strings.Contains(ipAddress, ".")
	isIPv6 := parsedIP != nil && strings.Contains(ipAddress, ":")
	if isIPv4 || isIPv6 {
		fmt.Printf("%v is a valid IP address\n", ipAddress)
		return true
	}

	fmt.Printf("%v is not a valid IP address\n", ipAddress)
	return false
}

func main() {
	randomString := RandomString(50)
	fmt.Println("Generated random string : ", randomString)
	fmt.Println("Generated random number : ", RandomNumber(16))

	fmt.Println("Generated random int in range : ", getRandomIntInRange(1, 3))

	freePort, err := getFreePort()
	if err != nil {
		fmt.Println("Error getting a free port")
		os.Exit(1)
	}
	fmt.Println("My free port is ", freePort)

	isValidIP("192.168.0.1")
	isValidIP("FE80:0000:0000:0000:0202:B3FF:FE1E:8329")
	isValidIP("FE80::0202:B3FF:FE1E:8329")
	isValidIP("FE80:")
	isValidIP("684D:1111:222:3333:4444:5555:6:77")
	isValidIP("http://[2001:db8:0:1]:80")
}
