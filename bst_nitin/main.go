package main

import (
	"fmt"
	"io"
	"os"
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

func inorderTraversal(n *node) {
	if n != nil {
		inorderTraversal(n.left)
		fmt.Print(n.data, " ")
		inorderTraversal(n.right)
	}
}
func preorderTraversal(n *node) {
	if n != nil {
		fmt.Print(n.data, " ")
		inorderTraversal(n.left)
		inorderTraversal(n.right)
	}
}
func postorderTraversal(n *node) {
	if n != nil {
		inorderTraversal(n.left)
		inorderTraversal(n.right)
		fmt.Print(n.data, " ")
	}
}

func min(n *node) int {
	if n == nil {
		return 0
	}
	for {
		if n.left == nil {
			return n.data
		}
		n = n.left
	}
}

func max(n *node) int {
	if n == nil {
		return 0
	}
	for {
		if n.right == nil {
			return n.data
		}
		n = n.right
	}
}

func levelOrderTraversal(n *node) {
	if n != nil {
		s := []*node{n}
		for len(s) > 0 {
			currentNode := s[0]
			if currentNode != nil {
				fmt.Printf(fmt.Sprint(currentNode.data) + " ")
				s = s[1:]
				if currentNode.left != nil {
					s = append(s, currentNode.left)
				}
				if currentNode.right != nil {
					s = append(s, currentNode.right)
				}
			}
		}
	}
}

func remove(n *node, key int) *node {
	if n == nil {
		return nil
	}
	if key < n.data {
		n.left = remove(n.left, key)
		return n
	}
	if key > n.data {
		n.right = remove(n.right, key)
		return n
	}
	if n.left == nil && n.right == nil {
		n = nil
		return nil
	}
	if n.left == nil {
		n = remove(n.right, key)
		return n
	}
	if n.right == nil {
		n = remove(n.left, key)
		return n
	}
	currentMostRightNode := n.right
	for currentMostRightNode != nil {
		if currentMostRightNode != nil && currentMostRightNode.left != nil {
			currentMostRightNode = currentMostRightNode.left
		} else {
			break
		}
	}
	n.data = currentMostRightNode.data
	n.right = remove(n.right, n.data)
	return n
}

func main() {
	fmt.Println("Binary Tree")
	var root *node
	root = insert(root, 20)
	root = insert(root, 15)
	root = insert(root, 25)
	root = insert(root, 18)
	root = insert(root, 10)
	root = insert(root, 30)
	root = insert(root, 21)

	fmt.Println("-----------printbst---------------")
	printbst(os.Stdout, root, 0, 'M')
	fmt.Println("------------stringify--------------")
	stringify(root, 0)
	//fmt.Println("-------------postorderTraversal-------------")
	//postorderTraversal(root)
	//fmt.Println("\n-----------preorderTraversal---------------")
	//preorderTraversal(root)
	//fmt.Println("\n------------inorderTraversal--------------")
	//inorderTraversal(root)
	//fmt.Println("\n------------max----------- : ", max(root))
	//fmt.Println("\n-----------min------------  :", min(root))
	//fmt.Println("\n------------levelOrderTraversal--------------")
	//levelOrderTraversal(root)
	//fmt.Println("\n------------search--------------", search(root, 1))
	//
	//root = insert(root, 21)
	//root = insert(root, 32)
	//fmt.Println("\n-----------stringify after duplicate insert---------------")
	//stringify(root, 0)
	//
	//remove(root, 32)
	//fmt.Println("\n-----------stringify after remove---------------")
	//stringify(root, 0)
	var stack Stack
	stack.roottoleafPrint(root)
}

type Stack []int

func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(str int) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (bool, int) {
	if s.isEmpty() {
		return true, 0
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return true, element
	}

}

func (s *Stack) roottoleafPrint(n *node) {

	if n == nil {
		return
	}
	s.Push(n.data)
	s.roottoleafPrint(n.left)
	if n.left == nil && n.right == nil {
		if !s.isEmpty() {
			fmt.Println(*s)
		}
	}
	s.roottoleafPrint(n.right)
	s.Pop()
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
