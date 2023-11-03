package main

// write a json string variable with demo data

var data = `{
    "id": 123456789,
    "owner": 123456789,
    "account_id": 123456789,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "mobile": "123456789",
    "edited": 123456789
}`

type User struct {
	ID        int    `json:"id"`
	Owner     int    `json:"owner"`
	AccountID int    `json:"account_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Edited    int    `json:"edited"`
}

var user = []User{
	{
		ID:        102,
		Owner:     568634,
		AccountID: 473223,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Mobile:    "123456789",
		Edited:    123456789,
	},
}

func main() {

}
