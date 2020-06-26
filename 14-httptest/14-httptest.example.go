package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/user/profile") && strings.HasSuffix(r.URL.Path, "/profile") {
			fmt.Fprintln(w, "Hello People profile, client")
		}
		if strings.HasPrefix(r.URL.Path, "/user/settings") && strings.HasSuffix(r.URL.Path, "/settings") {
			fmt.Fprintln(w, "I am the tags, client")
		}
	}))
	defer ts.Close()

	url := ts.URL + "/user/settings"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	greetings, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Running %s \n", url)
	fmt.Printf("Response body : %s \n", string(greetings))

	//time.Sleep(time.Minute * 2)
}
