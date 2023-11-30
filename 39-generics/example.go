package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return value, true
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop())  // Output: 3, true
	fmt.Println(intStack.Size()) // Output: 2

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack)         // Output: {[hello world]}
	fmt.Println(stringStack.Size())  // Output: 2
	fmt.Println(stringStack.data[0]) // Output: hello
	fmt.Println(stringStack.Pop())   // Output: world, true
	fmt.Println(stringStack.Size())  // Output: 1
}
