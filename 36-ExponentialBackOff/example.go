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

	// backoff.RetryNotify example
	var value []byte
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

	fmt.Println("RetryNotify")
	fmt.Println(string(value))

	// backoff.RetryNotifyWithData example
	operation := func() ([]byte, error) {
		resp, errResp := sendRequest()
		return resp, errResp
	}

	respData, err := backoff.RetryNotifyWithData(operation, exponentialBackOff, notify)
	if err != nil {
		fmt.Printf("error after retrying with RetryNotifyWithData: %v", err)
	}
	fmt.Println("data RetryNotifyWithData ")
	fmt.Println(string(respData))

	// backoff.RetryNotifyWithTimerAndData example
	/*operationWithTimerAndData := func() (io.ReadCloser, error) {
		resp, errResp := sendRequest()
		defer resp.Body.Close()
		return resp.Body, errResp
	}
	respDataWithTimer, err := backoff.RetryNotifyWithTimerAndData(operationWithTimerAndData, exponentialBackOff, notify, nil)
	if err != nil {
		fmt.Printf("error after retrying with RetryNotifyWithTimerAndData: %v", err)
	}

	data, err = io.ReadAll(respDataWithTimer)
	if err != nil {
		fmt.Printf("error reading data from RetryNotifyWithTimerAndData: %v", err)
	}
	fmt.Println(data)*/
}

func sendRequest() ([]byte, error) {
	requestURL := "https://jsonplaceholder.typicode.com/posts/1"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return []byte{}, err
	}
	defer res.Body.Close()

	// read data from response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
