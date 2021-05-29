package main

import (
	"fmt"
)

type node struct {
	value string
	left  *node
	right *node
}

func insert(n *node, v string) *node {
	if n == nil {
		return &node{v, nil, nil}
	} else if v <= n.value {
		n.left = insert(n.left, v)
	} else {
		n.right = insert(n.right, v)
	}
	return n
}

/* pre-oder DFS traversal */
func preorder(n *node) {
	if n != nil {
		fmt.Printf(n.value + " ")
		preorder(n.left)
		preorder(n.right)
	}
}

/* in-oder DFS traversal */
func inorder(n *node) {
	if n != nil {
		inorder(n.left)
		fmt.Printf(n.value + " ")
		inorder(n.right)
	}
}

/* post-oder DFS traversal */
func postorder(n *node) {
	if n != nil {
		postorder(n.left)
		postorder(n.right)
		fmt.Printf(n.value + " ")
	}
}

func max(n *node) string {
	if n == nil {
		return ""
	}
	for {
		if n.right == nil {
			return n.value
		}
		n = n.right
	}
}

func min(n *node) string {
	if n == nil {
		return ""
	}
	for {
		if n.left == nil {
			return n.value
		}
		n = n.left
	}
}

/* breadth first traversal */
func breadth(n *node) {
	if n != nil {
		s := []*node{n}
		for len(s) > 0 {
			currentNode := s[0]
			fmt.Printf(currentNode.value + " ")
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

func print(n *node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		print(n.left, level)
		fmt.Printf(format+"%s\n", n.value)
		print(n.right, level)
	}
}

func main() {
	var root *node
	root = insert(root, "20")
	root = insert(root, "18")
	root = insert(root, "22")
	root = insert(root, "11")
	root = insert(root, "19")
	root = insert(root, "21")
	root = insert(root, "25")

	print(root, 0)

	fmt.Println("Min : ", min(root))
	fmt.Println("Max : ", max(root))

	fmt.Println("Pre-order DFS traversal")
	preorder(root)
	fmt.Println("\nIn-order DFS traversal")
	inorder(root)
	fmt.Println("\nPost-order DFS traversal")
	postorder(root)
	fmt.Println("\nBFS traversal")
	breadth(root)
}

//==================================================================================

//package main
//
//import "fmt"
//
//type Node struct {
//	val   int
//	left  *Node
//	right *Node
//}
//
//func heightOfTree(root *Node) int {
//	if root == nil {
//		return 0
//	}
//
//	var left = heightOfTree(root.left)
//	var right = heightOfTree(root.right)
//
//	if left > right {
//		return left + 1
//	} else {
//		return right + 1
//	}
//}
//
//func printLevel(root *Node, h int) {
//	if root == nil {
//		return
//	}
//
//	if h == 1 {
//		fmt.Printf("%d ", root.val)
//		return
//	}
//	if h > 1 {
//		printLevel(root.left, h-1)
//
//		printLevel(root.right, h-1)
//	}
//}
//
//func levelOrder(root *Node) {
//	if root == nil {
//		return
//	}
//	var height = heightOfTree(root)
//	fmt.Printf("Height of a tree : %d \n", height)
//	fmt.Printf("Level Order Traversal : \n")
//
//	for i := 1; i <= height; i++ {
//		printLevel(root, i)
//	}
//
//	fmt.Printf("\n")
//}
//
//func main() {
//	// BST Representation
//	//         100
//	//        /   \
//	//      50    150
//	//     / \    /  \
//	//   20  80  110 200
//	//  /  \
//	// 10  30
//
//	root := Node{100, nil, nil}
//	root.left = &Node{50, nil, nil}
//	root.right = &Node{150, nil, nil}
//	root.left.left = &Node{20, nil, nil}
//	root.left.left.left = &Node{10, nil, nil}
//	root.left.left.right = &Node{30, nil, nil}
//	root.left.right = &Node{80, nil, nil}
//	root.right.left = &Node{110, nil, nil}
//	root.right.right = &Node{200, nil, nil}
//
//	levelOrder(&root)
//
//}
