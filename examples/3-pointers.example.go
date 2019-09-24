package examples

import(
	"fmt"
)

func main(){
	s := "first text"
	pointer := &s
	fmt.Println(s)
	fmt.Println(pointer)
	fmt.Println(*pointer)

}
