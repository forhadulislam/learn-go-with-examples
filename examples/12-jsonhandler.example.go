package main

import(
	"fmt"
	"encoding/json"
)

type Vehicle struct{
	Brand string
	Model string
	Year int
	Warranty int
}

func main(){

	// Marshaling based on a struct
	v := Vehicle{
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
	simpleJson := []byte(`{"Brand":"Hyundai","Model":"i30","Year":2017,"Warranty":5}`)
	vUnMarsh := &Vehicle{}
	error = json.Unmarshal(simpleJson, &vUnMarsh)
	fmt.Println( vUnMarsh )
	fmt.Printf( "This is a : %T \n", vUnMarsh )

}
