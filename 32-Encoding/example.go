package main

import (
	"fmt"
	"html"
	"net/url"
)

func main() {
	fmt.Println(html.EscapeString("<>"))
	fmt.Println(url.PathEscape("MY/VS"))
	fmt.Println(url.PathEscape("SOMETHING"))
	fmt.Println(url.PathEscape("A B"))
}
