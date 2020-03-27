package main

import (
	"fmt"
	"os"
	"text/template"
)

type Books struct {
	Title       string
	Author      string
	Description string
}

type BookList struct {
	Reader string
	List   []Books
}

func main() {

	// Example 1
	lib := Books{
		Title:       "Nothing to lose",
		Description: "You have nothing to lose",
	}

	t, err := template.New("todos").Parse("You have a book named '{{ .Title}}' with description: '{{ .Description}}'")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, lib)
	if err != nil {
		panic(err)
	}

	fmt.Println("")

	// Example 2
	paths := []string{
		"./resources/template.tmpl",
	}
	bookList := BookList{
		Reader: "John Doe",
		List: []Books{
			{
				Title:  "War and Peace",
				Author: " Leo Tolstoy",
				Description: `War and Peace broadly focuses on Napoleon's invasion of Russia in 1812 and follows three 
of the most well-known characters in literature: Prince Andrei Bolkonsky, 
who leaves his family behind to fight in the war against Napoleon; 
and Natasha Rostov, the young daughter of a nobleman who intrigues both men.`,
			},
			{
				Title:  "The Adventures of Huckleberry Finn",
				Author: "Mark Twain",
			},
			{
				Title:  "Hamlet ",
				Author: "William Shakespeare",
			},
		},
	}

	tpl, err := template.ParseFiles(paths...)
	err = tpl.Execute(os.Stdout, bookList)
	if err != nil {
		panic(err)
	}

}
