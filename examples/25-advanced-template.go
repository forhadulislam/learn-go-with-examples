package main

// https://www.calhoun.io/intro-to-templates-p3-functions/
// https://blog.golang.org/laws-of-reflection
// https://golang.org/pkg/reflect/
// https://golang.org/pkg/text/template/#FuncMap

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"reflect"
	"strconv"
	"text/template"
)

func main() {
	text := `
		{
			"_id": "5ebbb20e29d76a9f5217c5d0",
			"index": 0,
			"guid": "117ab7b9-c252-4272-b7ed-b4baa03131b5",
			"isActive": false,
			"balance": "$2,254.60",
			"picture": "http://placehold.it/32x32",
			"age": 26,
			"eyeColor": "green",
			"name": "Murray Estrada",
			"gender": "male",
			"company": "GEEKY",
			"email": "murrayestrada@geeky.com",
			"phone": "+1 (816) 491-2930",
			"address": "715 Alice Court, Jenkinsville, Mississippi, 2883",
			"about": "Consequat deserunt sit ex magna duis officia deserunt elit veniam. Anim excepteur sit adipisicing officia ullamco adipisicing voluptate nisi nostrud enim excepteur magna in. Dolor eiusmod ex labore aliqua irure elit. Qui proident sit veniam amet nostrud mollit mollit excepteur mollit veniam consequat. Ipsum aute voluptate minim nostrud exercitation ad sit nisi. Minim deserunt ex Lorem ea reprehenderit ut consequat. Ipsum eiusmod minim aliquip ex est ex incididunt veniam cupidatat.\r\n",
			"registered": "2020-02-07T08:02:59 -02:00",
			"latitude": 35.9457,
			"longitude": 133.234414,
			"tags": [
			  "est",
			  "anim",
			  "ad",
			  "qui",
			  "ipsum",
			  "eiusmod",
			  "dolore"
			],
			"friends": [
			  {
				"id": 0,
				"name": "Merle Rush"
			  },
			  {
				"id": 1,
				"name": "Hart Mayo"
			  },
			  {
				"id": 2,
				"name": "Sherry Gallagher"
			  }
			],
			"greeting": "Hello, Murray Estrada! You have 3 unread messages.",
			"favoriteFruit": "apple",
			"moredata": [
				{"post": "I have"},
				{"post": "nothing"},
				{"post": "to"},
				{"post": [0, 1, 1, 3, 2]},
				{"post": "do"},
				{"post": "now!"},
				99,
				["kla", "kli", "klu"]
			]
		  }
	`

	user := make(map[string]interface{})
	if err := json.Unmarshal([]byte(text), &user); err != nil {
		panic(err)
	}

	tf := template.FuncMap{
		"getAgeCategory": func(i interface{}) string {
			v := fmt.Sprintf("%v", i)
			iv, err := strconv.Atoi(v)
			if err != nil {
				return ""
			}
			if iv <= 20 {
				return "Young duck"
			}else if iv <= 40 {
				return "Happy hunk"
			}else{
				return "Relaxed punk"
			}
		},
		"isInt": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
				return true
			default:
				return false
			}
		},
		"isString": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.String:
				return true
			default:
				return false
			}
		},
		"isSlice": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Slice:
				return true
			default:
				return false
			}
		},
		"isArray": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Array:
				return true
			default:
				return false
			}
		},
		"isMap": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Map:
				return true
			default:
				return false
			}
		},
		"concat": func(i interface{}, j interface{}) string {
			str := fmt.Sprintf("%s%s", i, j)
			return str
		},
		"nullVal": func(i interface{}, def string) string {
			if i != nil {
				return fmt.Sprintf("%s", i)
			}
			return def
		},
		"emptyString": func(input string, def string) string {
			if input == "" {
				return def
			}
			return input
		},
	}

	paths := []string{
		"./resources/profile.tmpl",
	}

	name := path.Base(paths[0])
	t, err := template.New(name).Funcs(tf).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	if err = t.Execute(os.Stdout, &user); err != nil {
		panic(err)
	}
}
