package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "<h2>Welcome to the home Page</h2>")
	fmt.Fprintf(w, "%s", "How are you doing today?")
}

func work(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[ %s", time.Now())
	fmt.Fprintf(w, "]")
}

func main() {
	response, _ := http.Get("https://golangcode.com/")

	// The line below would fail because Body = io.ReadCloser
	// fmt.Printf(response.Body)

	// ...so we convert it to a string by passing it through
	// a buffer first. A 'costly' but useful process.
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	defer response.Body.Close()

	serverPort := "8085"
	http.HandleFunc("/home", home)
	http.HandleFunc("/workTime", work)
	fmt.Println("Server is running at port : " + serverPort)
	http.ListenAndServe(":"+serverPort, nil)
}
