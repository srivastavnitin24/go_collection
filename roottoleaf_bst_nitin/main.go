package main

import (
	"fmt"
	"strconv"
)

type node struct {
	val   int
	left  *node
	right *node
}

func main() {
	var root *node
	root = insert(root, 20)
	root = insert(root, 15)
	root = insert(root, 25)
	root = insert(root, 10)
	root = insert(root, 28)
	root = insert(root, 21)
	root = insert(root, 30)

	printTree(root, 0)

	arr := rootToLeaf(root)
	for i, s := range arr {
		fmt.Println("\n", i, " : ", s)
	}

}

func rootToLeaf(n *node) []string {
	var s []string
	if n == nil {
		return s
	}
	helper(n, strconv.Itoa(n.val), &s)
	return s
}

func helper(n *node, path string, s *[]string) {
	if n.left == nil && n.right == nil {
		*s = append(*s, path)
	}
	if n.left != nil {
		helper(n.left, path+" -> "+strconv.Itoa(n.left.val), s)
	}
	if n.right != nil {
		helper(n.right, path+" -> "+strconv.Itoa(n.right.val), s)
	}

}

func printTree(n *node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "	"
		}
		level++
		format += " ----[ "

		printTree(n.right, level)
		fmt.Printf(format+" %s\n ", fmt.Sprint(n.val))
		printTree(n.left, level)
	}
}

func insert(n *node, key int) *node {
	if n == nil {
		return &node{val: key}
	}
	if key <= n.val {
		n.left = insert(n.left, key)
		return n
	} else {
		n.right = insert(n.right, key)
		return n
	}
}
