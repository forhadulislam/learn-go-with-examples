package main

import(
	"fmt"
	"encoding/json"
)

type Transport struct{
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year int `json:"year"`
	Warranty int `json:"warranty"`
}

func main(){

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
	if (error != nil){
		fmt.Errorf("Could not Marshal the struct! Error: %v", error)
	}
	fmt.Println( string(vMarsh) )
	fmt.Printf( "This is a : %T \n", vMarsh )

	// 			UnMarshaling
	// UnMarshaling based on a struct
	simpleJson := []byte(`{"Brand":"Hyundai","Model":"i30","Warranty":5,"Year":2017}`)
	vUnMarsh := &Transport{}
	error = json.Unmarshal(simpleJson, &vUnMarsh)
	fmt.Println( vUnMarsh )
	fmt.Printf( "This is a : %T \n", vUnMarsh )

}
