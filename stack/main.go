package main

import "fmt"

type Stack []string

func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (bool, string) {
	if s.isEmpty() {
		return true, ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return true, element
	}

}

func main() {
	var stack Stack
	stack.Push("India")
	stack.Push("Love")
	stack.Push("I")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if x {
			fmt.Println("Poped : ", y)
		}
	}
}
