package main

import (
	"fmt"
	"regexp"
)

func main() {
	// pattern example #1
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	// Pattern #2
	founda, err := regexp.MatchString("g([a-z]+)ng", "golang")
	fmt.Printf("found=%v, err=%v \n", founda, err)

	// Pattern #3
	foundb, err := regexp.MatchString("\"([a-z]+)\"", "\"anythin\"")
	fmt.Printf("found=%v, err=%v \n", foundb, err)

	// Pattern #4 JSON
	re = regexp.MustCompile("\"title\": \"")
	stringC := re.ReplaceAllString(`{
			  "data": [{
				"type": "articles",
				"id": "1",
				"attributes": {
				  "title": "JSON:API paints my bikeshed!",
				  "body": "The shortest article. Ever.",
				  "created": "2015-05-22T14:56:29.000Z",
				  "updated": "2015-05-22T14:56:28.000Z"
				},
				"relationships": {
				  "author": {
					"data": {"id": "42", "type": "people"}
				  }
				}
			  }],
			  "included": [
				{
				  "type": "people",
				  "id": "42",
				  "attributes": {
					"name": "John",
					"age": 80,
					"gender": "male"
				  }
				}
			  ]
			}`,
		"IHAVEREPLACETIT")
	fmt.Printf("stringC = %v n", stringC)

	// Pattern #5
	var Myregex = `//==start==
		.*
		//==end==`

	var originalString = `		
		.myClass1 {
			color: green;
		}		
		.myClass2 {
			color: white;
		}

		//==start==
		text
		//==end==`
	re = regexp.MustCompile(Myregex)
	s := re.ReplaceAllString(originalString, "replaced")
	fmt.Println(s)

	// Pattern #5
	re = regexp.MustCompile(`(^|[^_])\bproducts\b([^_]|$)`)
	var sample = `{% macro products_list(products) %}
					{% for product in products %}
						productsList
					{% endfor %}
				{% endmacro %}`
	s = re.ReplaceAllString(sample, `$1.$2`)
	fmt.Println(s)
}
