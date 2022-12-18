package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
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

var ret int

func diameterOfBinaryTree(root *node) int {
	ret = 0
	helper(root)
	if ret == 0 {
		return ret
	}
	return ret - 1
}

func helper(root *node) int {
	if root == nil {
		return 0
	}
	left := helper(root.left)
	right := helper(root.right)
	if ret < left+right+1 {
		ret = left + right + 1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	fmt.Println("Binary Tree")
	var root *node
	root = insert(root, 1)
	root = insert(root, 2)
	//root = insert(root, 3)
	//root = insert(root, 4)
	//root = insert(root, 5)

	fmt.Println("------------stringify--------------")
	stringify(root, 0)

	diameter := diameterOfBinaryTree(root)
	fmt.Println("------------diameterOfBinaryTree--------------", diameter)

}
