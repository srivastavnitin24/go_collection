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

func heightOfBst(n *node) int {
	if n == nil {
		return 0
	}

	finalHeight := 0
	leftHeight := heightOfBst(n.left)
	rightHeight := heightOfBst(n.right)

	if leftHeight > rightHeight {
		finalHeight = leftHeight + 1
	} else {
		finalHeight = rightHeight + 1
	}
	return finalHeight
}

func diameterOfBst(n *node) int {

	if n == nil {
		return 0
	}
	leftHeight := heightOfBst(n.left)
	rightHeight := heightOfBst(n.right)

	leftDiameter := diameterOfBst(n.left)
	rightDiameter := diameterOfBst(n.right)

	return max(leftHeight+rightHeight+1, max(leftDiameter, rightDiameter))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Binary Tree")
	var root *node
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
	root.right = &node{data: 3}
	root.left.right = &node{data: -4}
	root.right.left = &node{data: 5}
	root.right.right = &node{data: 6}
	root.right.left.left = &node{data: 7}
	root.right.left.right = &node{data: 8}

	//fmt.Println("-----------printbst---------------")
	//printbst(os.Stdout, root, 0, 'M')
	fmt.Println("------------stringify--------------")
	stringify(root, 0)

	fmt.Println("\n------------heightOfBst--------------", heightOfBst(root))

	fmt.Println("\n------------diameterOfBst--------------", diameterOfBst(root))

}
