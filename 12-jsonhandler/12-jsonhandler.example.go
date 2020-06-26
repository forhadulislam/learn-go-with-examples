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
}

func main() {

	// Marshaling based on a struct
	v := Transport{
		Brand:    "Hyundai",
		Model:    "i30",
		Year:     2017,
		Warranty: 5,
	}

	// 			Marshaling
	// Marshaling to json byte slice
	vMarsh, error := json.Marshal(v)
	if error != nil {
		fmt.Errorf("could not Marshal the struct! Error: %v", error)
	}
	fmt.Println(string(vMarsh))
	fmt.Printf("This is a : %T \n", vMarsh)

	// Marshaling empty structs
	vEmpty := Transport{}
	vEmptyMarsh, err := json.Marshal(vEmpty)
	if err != nil {
		fmt.Errorf("could not Marshal the struct! Error: %v", error)
	}
	fmt.Println(string(vEmptyMarsh))
	fmt.Printf("This is a : %T \n", vEmptyMarsh)

	// 			UnMarshaling
	// UnMarshaling based on a struct
	simpleJson := []byte(`{"Brand":"Hyundai","Model":"i30","Warranty":5,"Year":2017}`)
	vUnMarsh := &Transport{}
	error = json.Unmarshal(simpleJson, &vUnMarsh)
	fmt.Println(vUnMarsh)
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
	error = json.Unmarshal([]byte(jsonString), &jBigMarsh)
	fmt.Println(jBigMarsh)
	for dt := range jBigMarsh{
		fmt.Println(dt)
		fmt.Println(jBigMarsh[dt]["endpoint"])
		if jBigMarsh[dt]["details"] == nil {
			fmt.Println("This is nil {}")
		}else{
			vMarsh, err := json.Marshal(jBigMarsh[dt]["details"])
			if err != nil {
				fmt.Errorf("could not Marshal the struct! Error: %v", error)
			}
			fmt.Println(string(vMarsh))
		}
		fmt.Println(jBigMarsh[dt]["details"])
	}
	fmt.Printf("This is a : %T \n", jBigMarsh)

}
