package main

import (
	"fmt"
	"strings"
)

func main() {
var b strings.Builder
	b.Grow(32)
	for i, p := range []int{2, 3, 5, 7, 11, 13} {
		fmt.Fprintf(&b, "%d:%d, ", i+1, p)
	}
	s := b.String()   // no copying
	s = s[:b.Len()-2] // no copying (removes trailing ", ")
	fmt.Println(s)
}
