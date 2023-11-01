package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	exponentialBackOff := backoff.NewExponentialBackOff()
	exponentialBackOff.MaxElapsedTime = 3 * time.Minute

	var value *http.Response
	var err error

	retryable := func() error {
		value, err = sendRequest()
		return err
	}

	notify := func(err error, t time.Duration) {
		log.Printf("error: %v happened at time: %v", err, t)
	}

	err = backoff.RetryNotify(retryable, exponentialBackOff, notify)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}

	fmt.Println(value.StatusCode)

}

func sendRequest() (*http.Response, error) {
	requestURL := "https://jsonplaceholder.typicode.com/posts/1"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return &http.Response{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return &http.Response{}, err
	}

	// read data from response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &http.Response{}, err
	}

	fmt.Println(string(body))

	return res, nil
}
