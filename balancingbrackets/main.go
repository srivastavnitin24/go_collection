package main

//
//import "fmt"
//
//type stack []rune
//
//func (s *stack) push(ch rune) {
//	fmt.Println("======= : ", ch)
//	*s = append(*s, ch)
//}
//
//func (s *stack) pop() (rune, bool) {
//	if s.empty() {
//		return , false
//	}
//	index := len(*s) - 1
//	popElement := (*s)[index]
//	*s = (*s)[:index]
//	return popElement, true
//}
//
//func (s *stack) top() rune {
//	return (*s)[len(*s)-1]
//}
//
//func (s *stack) empty() bool {
//	if len(*s) == 0 {
//		return true
//	}
//	return false
//}
//
//func main() {
//	var s *stack
//	fmt.Println("Are the brackets in correct order : ", s.isBalanced("[{}(){()}]"))
//}
//
//func (s *stack) isBalanced(expr string) bool {
//	for i := 0; i < len(expr); i++ {
//		//expr := string(expr[i])
//		expr := rune(expr[i])
//		if expr == '(' || expr == '[' || expr == '{' { //when it is opening bracket, push into stack
//			s.push(expr)
//			continue
//		}
//
//		if s.empty() { //stack cannot be empty as it is not opening bracket, there must be closing bracket
//			return false
//		}
//
//		switch *(expr)[i] {
//		case ')': //for closing parenthesis, pop it and check for braces and square brackets
//			ch := s.top()
//			s.pop()
//			if ch == '{' || ch == "[" {
//				return false
//			}
//			break
//		case '}': //for closing braces, pop it and check for parenthesis and square brackets
//			ch := s.top()
//			s.pop()
//			if ch == "(" || ch == "[" {
//				return false
//			}
//			break
//		case ']': //for closing square bracket, pop it and check for braces and parenthesis
//			ch := s.top()
//			s.pop()
//			if ch == "(" || ch == "{" {
//				return false
//			}
//			break
//		}
//	}
//	return s.empty()
//}
