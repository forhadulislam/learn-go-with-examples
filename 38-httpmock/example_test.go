package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestFetchArticles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// our database of articles
	users := make([]User, 0)

	// mock to list out the articles
	httpmock.RegisterResponder("GET", "https://api.gitty.com/users",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, user)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		})

	// return an article related to the request with the help of regexp submatch (\d+)
	httpmock.RegisterResponder("GET", `=~^https://api\.gitty\.com/users/id/(\d+)\z`,
		func(req *http.Request) (*http.Response, error) {
			// Get ID from request
			id := httpmock.MustGetSubmatchAsUint(req, 1) // 1=first regexp submatch
			return httpmock.NewJsonResponse(200, map[string]interface{}{
				"id":   id,
				"name": "My Great Article",
			})
		})

	// mock to add a new article
	httpmock.RegisterResponder("POST", "https://api.gitty.com/users",
		func(req *http.Request) (*http.Response, error) {
			userData := User{}
			if err := json.NewDecoder(req.Body).Decode(&userData); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}

			users = append(users, userData)

			resp, err := httpmock.NewJsonResponse(200, userData)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		})

	// mock to add a specific article, send a Bad Request response
	// when the request body contains `"type":"toy"`
	httpmock.RegisterMatcherResponder("POST", "https://api.gitty.com/users",
		httpmock.BodyContainsString(`"type":"toy"`),
		httpmock.NewStringResponder(400, `{"reason":"Invalid article type"}`))

	// do stuff that adds and checks articles

	// fetch all users
	resp, err := sendHTTPRequest(http.MethodGet, "https://api.gitty.com/users", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)

	resp, err = sendHTTPRequest(http.MethodGet, "https://api.gitty.com/users/id/1", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}

func sendHTTPRequest(httpMethod, requestURL string, data io.ReadCloser) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, requestURL, data)
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
