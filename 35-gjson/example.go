package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

// write a json string variable with demo data

var data = `{
    "id": 123456789,
    "owner": 123456789,
    "account_id": 123456789,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "mobile": "123456789",
    "edited": 123456789,
    "user_details": {
        "first_name": "John",
        "last_name": "Doe",
        "email": "",
        "mobile": "123456789",
        "phone": "",
        "address": {
            "street": "Street 1",
            "city": "City",
            "country": "Country",
            "postalcode": "12345"
        },
        "company": {
            "name": "Company",
            "business_id": "1234567-8",
            "address": {
                "street": "Street 1",
                "city": "City",
                "country": "Country",
                "postalcode": "12345"
            }
        }
    }
}`

func main() {

	s := gjson.GetBytes([]byte(data), "name")
	fmt.Println(s)
	fmt.Println(s.Get("ids").Int())
	fmt.Println(s.Get("ids").Array())
	fmt.Println(s.Get("ids").Map())
	fmt.Println(gjson.Get(data, "user_details").Map())

	a := gjson.Parse(data)
	fmt.Println(a.Get("user_details.address"))
	fmt.Println(a.Get("user_details").Get("company"))
}
