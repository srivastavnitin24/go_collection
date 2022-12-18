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

var ret int

func BinaryTreeDiameter(root *node) int {
	if root == nil {
		return 0
	}
	left := BinaryTreeDiameter(root.left)
	right := BinaryTreeDiameter(root.right)
	if ret < left+right+1 {
		ret = left + right + 1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func BinaryTreeDiameterForward(root *node) int {
	ret = 0
	BinaryTreeDiameter(root)
	if ret == 0 {
		return ret
	}
	return ret - 1
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

	/* Construct the following tree
	       1
	     /   \
	    /     \
	   2       3
	    \     / \
	    -4   5   6
	        / \
	       7   8
	*/

	root = &node{data: 1}
	root.left = &node{data: 2}
	root.right = &node{data: 3}
	root.left.right = &node{data: -4}
	root.right.left = &node{data: 5}
	root.right.right = &node{data: 6}
	root.right.left.left = &node{data: 7}
	root.right.left.right = &node{data: 8}

	fmt.Println("------------stringify--------------")
	stringify(root, 0)

	fmt.Println("------------Binary Tree Diameter--------------", BinaryTreeDiameterForward(root))
	fmt.Println("------------find Max Sum Path--------------", maxPathSum(root))
}

var ans int

func maxPathSum(root *node) int {
	ans = math.MinInt32
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return root.data
	}
	maxPath(root)
	return ans
}
func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxPath(root *node) int {
	if root == nil {
		return 0
	}
	left := mymax(0, maxPath(root.left))
	right := mymax(0, maxPath(root.right))
	sum := left + right + root.data
	ans = mymax(ans, sum)
	return mymax(left, right) + root.data

}
