package main

import (
	"fmt"
	"io"
)

type node struct {
	data  int
	left  *node
	right *node
}

func insert(n *node, key int) *node {
	if n == nil {
		return &node{data: key}
	} else if !search(n, key) {
		if key <= n.data {
			n.left = insert(n.left, key)
		} else {
			n.right = insert(n.right, key)
		}
	} else {
		fmt.Printf("\n Cannot insert duplicate key : %d  as its already in the tree", key)
	}
	return n

}

func search(n *node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.data {
		return search(n.left, key)
	} else if key > n.data {
		return search(n.right, key)
	}
	return true

}

func printbst(w io.Writer, n *node, ns int, ch rune) {
	if n != nil {
		for i := 0; i < ns; i++ {
			fmt.Fprintf(w, " ")
		}
		fmt.Fprintf(w, "%c:%v\n", ch, n.data)
		printbst(w, n.left, ns+2, 'L')
		printbst(w, n.right, ns+2, 'R')
	}
}

func stringify(n *node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "	"
		}
		format += "-----["
		level++

		stringify(n.right, level)

		fmt.Printf(format+"%s\n", fmt.Sprint(n.data))

		stringify(n.left, level)
	}
}

type Stack []int

func (s *Stack) push(key int) {
	*s = append(*s, key)
}

func (s *Stack) pop() (bool, int) {
	if s.isEmpty() {
		return true, 0
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return true, element
	}

}

func (s *Stack) printstack() {
	if s.isEmpty() {
		return
	} else {
		index := len(*s)
		*s = (*s)[:index]
		fmt.Println(*s)
		return
	}

}

func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Binary Tree")
	var root *node
	var st Stack
	//root = insert(root, 20)
	//root = insert(root, 15)
	//root = insert(root, 25)
	//root = insert(root, 18)
	//root = insert(root, 10)
	//root = insert(root, 30)
	//root = insert(root, 21)
	//root = insert(root, 21)

	root = &node{data: 1}
	root.left = &node{data: 2}
	root.left.left = &node{data: 4}
	root.left.right = &node{data: 5}
	root.left.left.left = &node{data: 8}
	root.left.left.right = &node{data: 9}
	root.right = &node{data: 3}

	root.right.left = &node{data: 6}
	root.right.right = &node{data: 7}
	root.right.left.left = &node{data: 10}
	root.right.left.right = &node{data: 11}
	root.right.right.left = &node{data: 12}
	root.right.right.left.right = &node{data: -3}

	//fmt.Println("-----------printbst---------------")
	//printbst(os.Stdout, root, 0, 'M')
	fmt.Println("------------stringify--------------")
	stringify(root, 0)

	st.printallpathofsumk(root)
	st.printallpathfromroottoleaf(root)

}

var sum = 0
var Ksum = 20

func (s *Stack) printallpathofsumk(n *node) {
	if n == nil {
		return
	}
	sum = sum + n.data
	s.push(n.data)
	if sum == Ksum {
		fmt.Println("printallpathofsumk : ", *s, sum)
	}
	s.printallpathofsumk(n.left)

	s.printallpathofsumk(n.right)
	sum = sum - n.data
	s.pop()
}

var IndividualSum = 0

func (s *Stack) printallpathfromroottoleaf(n *node) {
	if n == nil {
		return
	}
	IndividualSum = IndividualSum + n.data
	s.push(n.data)
	if n.left == nil && n.right == nil {
		fmt.Println("printallpathfromroottoleaf : ", *s, " and sum : ", IndividualSum)
	}
	s.printallpathfromroottoleaf(n.left)
	s.printallpathfromroottoleaf(n.right)
	IndividualSum = IndividualSum - n.data
	s.pop()

}
