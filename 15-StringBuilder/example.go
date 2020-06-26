package main

import (
	"fmt"
	"strings"
)

type Data struct {
	Name  string
	Score int
}

func main() {
	d := []Data{
		{
			Name:  "John",
			Score: 10,
		},
		{
			Name:  "Sid",
			Score: 7,
		},
		{
			Name:  "Noman",
			Score: 12,
		},
		{
			Name:  "Meew",
			Score: 11,
		},
		{
			Name:  "Thaner",
			Score: 19,
		},
	}
	var output strings.Builder
	//output.Grow(32)
	for index, u := range d {
		if u.Score%2 == 0 {
			output.WriteString(u.Name + " is even")
		} else {
			output.WriteString(u.Name + " is odd")
		}
		if index >= 0 && index < (len(d)-1) {
			output.WriteString(",")
		} else {
			output.WriteString(".")
		}
	}
	fmt.Println(output.String())
}
