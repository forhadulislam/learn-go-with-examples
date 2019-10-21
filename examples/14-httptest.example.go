package main

import (
	"fmt"
	"strings"
	"time"
	"net/http"
	"net/http/httptest"
	"log"
	"io/ioutil"
)

func home( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "%s", "<h2>Welcome to the home Page</h2>" )
	fmt.Fprintf(w, "%s", "How are you doing today?" )
}

func work( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "[ %s", time.Now() )
	fmt.Fprintf(w, "]")
}

func main() {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/configurations/plans") && strings.HasSuffix(r.URL.Path, "/plans") {
			fmt.Fprintln(w, "Hello plans, client")
		}
		if strings.HasPrefix(r.URL.Path, "/configurations/plans") && strings.HasSuffix(r.URL.Path, "/tags") {
			fmt.Fprintln(w, "I am the tags, client")
		}
	}))
	defer ts.Close()

	//url := ts.URL + "/configurations/plans/sranPlan/tags"
	url := ts.URL + "/configurations/plans"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s \n", url)
	fmt.Printf("%s", greeting)

}
