package main

import (
	"fmt"
	"math"
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

func findMaxPathSum(node *node, maxSum int) int {

	if node == nil {
		return 0
	}

	left := findMaxPathSum(node.left, maxSum)

	right := findMaxPathSum(node.right, maxSum)

	max := maxSum
	max = MAX(max, node.data)
	max = MAX(max, node.data+left)
	max = MAX(max, node.data+right)
	max = MAX(max, node.data+left+right)

	return MAX(node.data, node.data+MAX(left, right))
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	var root *node
	//root = insert(root, 1)
	//root = insert(root, 2)
	//root = insert(root, 10)
	//root = insert(root, -1)
	//root = insert(root, -4)
	//root = insert(root, -5)
	//root = insert(root, -6)
	//root = insert(root, 4)
	//root = insert(root, 7)
	//root = insert(root, 4)
	//root = insert(root, -2)

	root = &node{data: 1}
	root.left = &node{data: 10}
	root.right = &node{data: 34}
	root.left.left = &node{data: 45}
	root.left.right = &node{data: -9}
	root.right.left = &node{data: -67}
	root.right.right = &node{data: 100}

	fmt.Println("------------stringify--------------")
	stringify(root, 0)

	fmt.Println("------------stringify--------------", findMaxPathSum(root, math.MinInt32))

}
