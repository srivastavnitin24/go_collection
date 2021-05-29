package main

import (
	"fmt"
	"io"
	"os"
)

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

//Tree
func (t *Tree) insertTree(key int) *Tree {
	if t.root == nil {
		t.root = &Node{data: key}
	} else {
		t.root.insertNode(key)
	}
	return t
}

//Node
func (n *Node) insertNode(key int) {
	if n == nil {
		return
	} else if key <= n.data {
		if n.left == nil {
			n.left = &Node{data: key}
		} else {
			n.left.insertNode(key)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: key}
		} else {
			n.right.insertNode(key)
		}
	}
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.data)
		stringify(n.right, level)
	}
}

func printTree(w io.Writer, n *Node, ns int, ch rune) {
	if n == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, n.data)
	printTree(w, n.left, ns+2, 'L')
	printTree(w, n.right, ns+2, 'R')
}

func inorderTraversal(n *Node) {
	if n != nil {
		inorderTraversal(n.left)
		fmt.Print(n.data, " : ")
		inorderTraversal(n.right)
	}
}

func preorderTraversal(n *Node) {
	if n != nil {
		fmt.Print(n.data, " : ")
		preorderTraversal(n.left)
		preorderTraversal(n.right)
	}
}

func postorderTraversal(n *Node) {
	if n != nil {
		postorderTraversal(n.left)
		postorderTraversal(n.right)
		fmt.Print(n.data, " : ")
	}
}

func search(n *Node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.data {
		return search(n.left, key)
	}
	if key > n.data {
		return search(n.right, key)
	}
	return true
}

func Min(n *Node) int {
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

func Max(n *Node) int {
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

// internal recursive function to remove an item
func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.data {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.data {
		node.right = remove(node.right, key)
		return node
	}
	// key == node.data
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.data = leftmostrightside.data
	node.right = remove(node.right, node.data)
	return node
}

func breadthTraversal(n *Node) {
	fmt.Println(" breadthTraversal ------- ")
	if n != nil {
		s := []*Node{n}
		for len(s) > 0 {
			currentNode := s[0]
			fmt.Print(fmt.Sprint(currentNode.data) + "  ")
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

func main() {
	fmt.Println("Binary Tree")
	tree := &Tree{}
	tree.insertTree(20)
	tree.insertTree(18)
	tree.insertTree(22)
	tree.insertTree(11)
	tree.insertTree(19)
	tree.insertTree(21)
	tree.insertTree(25)
	printTree(os.Stdout, tree.root, 0, 'M')

	breadthTraversal(tree.root)
	fmt.Println()
	// t.insertTree(4).insertTree(41).insertTree(45).insertTree(34).insertTree(49).insertTree(55).insertTree(5).insertTree(98)
	//stringify(t.root, 0)

	inorderTraversal(tree.root)
	fmt.Println()
	preorderTraversal(tree.root)
	fmt.Println()
	postorderTraversal(tree.root)

	//search
	fmt.Println()
	flag := search(tree.root, 0)
	fmt.Println("flag for search : ", flag)

	fmt.Println()
	fmt.Println("Max: ", Max(tree.root))

	fmt.Println()
	fmt.Println("Min: ", Min(tree.root))

	n := remove(tree.root, 9)
	fmt.Println("-------- :  ", n.data)

	inorderTraversal(tree.root)
	fmt.Println()
}
