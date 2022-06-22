package main

import (
	"encoding/json"
	"fmt"
)

type Transport struct {
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Year     int    `json:"year"`
	Warranty int    `json:"warranty"`
	hidden   string `json:"hidden"`
}

func main() {

	// Marshaling based on a struct
	v := Transport{
		Brand:    "Hyundai",
		Model:    "i30",
		Year:     2017,
		Warranty: 5,
		hidden:   "some hidden data",
	}

	// Marshaling to json byte slice
	vMarsh, err := json.Marshal(v)
	if err != nil {
		fmt.Println(fmt.Errorf("could not Marshal the struct! Error: %v", err))
	}
	fmt.Println(string(vMarsh))
	fmt.Printf("This is a : %T \n", vMarsh)

	// Marshaling empty structs
	vEmpty := Transport{}
	vEmptyMarsh, err := json.Marshal(vEmpty)
	if err != nil {
		fmt.Println(fmt.Errorf("could not Marshal the struct! Error: %v", err))
	}
	fmt.Println(string(vEmptyMarsh))
	fmt.Printf("This is a : %T \n", vEmptyMarsh)

	// UnMarshaling based on a struct
	simpleJson := []byte(`{"Brand":"Hyundai","Model":"i30","Warranty":5,"Year":2017, "hidden": "my hidden data"}`)
	vUnMarsh := &Transport{}
	err = json.Unmarshal(simpleJson, &vUnMarsh)
	fmt.Println(vUnMarsh)
	fmt.Printf("hidden data is empty: %t \n", vUnMarsh.hidden == "")
	fmt.Printf("This is a : %T \n", vUnMarsh)

	// Big string
	jsonString := `{
	  "userStatus": {
		"endpoint": "/user/login",
		"data": {}
	  },
	  "userLogs": {
		"endpoint": "/user/logs",
		"data": {}
	  },
	  "info": {
		"endpoint": "/user/info",
		"details": {
		  "activities": [
			{
			  "title": "login",
			  "uploadSubType": "321",
			  "location": {
				"latitude": "20.009",
				"longitude": "90.09"
			  }
			},
			{
			  "title": "register",
			  "uploadSubType": "322",
			  "location": {
				"latitude": "20.009",
				"longitude": "90.09"
			  }
			}
		  ]
		}
	  }
	}`
	jBigMarsh := map[string]map[string]interface{}{}
	//jBigMarsh := map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonString), &jBigMarsh)
	fmt.Println(jBigMarsh)
	for dt := range jBigMarsh {
		fmt.Println(dt)
		fmt.Println(jBigMarsh[dt]["endpoint"])
		if jBigMarsh[dt]["details"] == nil {
			fmt.Println("This is nil {}")
		} else {
			vMarsh, err := json.Marshal(jBigMarsh[dt]["details"])
			if err != nil {
				fmt.Errorf("could not Marshal the struct! Error: %v", err)
			}
			fmt.Println(string(vMarsh))
		}
		fmt.Println(jBigMarsh[dt]["details"])
	}
	fmt.Printf("This is a : %T \n", jBigMarsh)

	// UnMarshaling based on a struct
	type LogsData struct {
		Url string `json:"url"`
	}
	fmt.Println("UnMarshaling based on a struct with additional data")
	jsonWithUnnecessaryData := []byte(`
			[
				{"url":"https://url.1/","model":"i30"},
				{"url":"https://url.2/","model":"i30"},
				{"url":"https://url.3/","model":"i30"},
				{"model":"i30"}
			]			
		`)
	var vUnMarsh2 []LogsData
	err = json.Unmarshal(jsonWithUnnecessaryData, &vUnMarsh2)
	fmt.Println(err)
	fmt.Println(vUnMarsh2)
	fmt.Println(len(vUnMarsh2))
	fmt.Println(vUnMarsh2[3].Url)
	fmt.Printf("This is a : %T \n", vUnMarsh2)

	// UnMarshaling for a json.Rawmessage
	type LogsData2 struct {
		Url json.RawMessage `json:"url"`
	}
	fmt.Println("UnMarshaling for a json.Rawmessage")
	jsonWithUnnecessaryData = json.RawMessage(`
			[
				{"url":"https://url.1/","model":"i30"},
				{"url":"https://url.2/","model":"i30"},
				{"url":"https://url.3/","model":"i30"},
				{"model":"i30"}
			]			
		`)
	var vUnMarsh3 []LogsData2
	err = json.Unmarshal(jsonWithUnnecessaryData, &vUnMarsh3)
	fmt.Println(err)
	fmt.Println(vUnMarsh3)
	fmt.Println(len(vUnMarsh3))
	fmt.Println(string(vUnMarsh3[2].Url))
	fmt.Printf("This is a : %T \n", vUnMarsh3)

}
