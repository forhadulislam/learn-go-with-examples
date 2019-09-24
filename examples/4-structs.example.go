package example

import(
	"fmt"
)

type Vehicle struct{
	Brand string
	Model string
	Year int
	Warranty int
}

func main() {
	v := Vehicle{
		Brand:    "Hyundai",
		Model:    "i30",
		Year:     2019,
		Warranty: 10,
	}

	fmt.Println(v)

	annonymusStruct := struct {
		Name string
		Email string
		Phone string
	}{
		"John",
		"john.doe@email.com",
		"+187778AAA",
	}

	fmt.Println(annonymusStruct)
}