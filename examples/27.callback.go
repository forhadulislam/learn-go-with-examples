package main

import "strings"

func callBack(sentence string, function func(callBackString string))  {
	println("Input: ", sentence)
	function(sentence)
	println("I am a callback and I am executed!")
}

func main() {
	callBack("Hello call back", func(s string) {
		b := strings.Split(s, " ")
		println( len(s) )
		println( len(b) )
	})
}
