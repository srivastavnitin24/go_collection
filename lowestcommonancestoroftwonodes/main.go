package main

import (
	"fmt"
	"strings"
)

type hierarchical interface {
	lowestCommonManager(*node, string, string) *node
	search(*node, string, string)
}

type node struct {
	employee string
	left     *node
	right    *node
}

func main() {

	var root *node
	root = &node{employee: "CEO-Claire"}
	root.left = &node{employee: "Nitin"}
	root.right = &node{employee: "AMIT"}
	root.left.right = &node{employee: "JONES"}
	root.left.left = &node{employee: "krishna"}
	root.right.left = &node{employee: "MARK"}
	root.right.right = &node{employee: "SUMIT"}
	root.right.left.left = &node{employee: "Dan"}
	root.right.left.right = &node{employee: "JEFF"}
	root.right.right.left = &node{employee: "WRIGHT"}
	root.right.right.right = &node{employee: "ANJELINA"}
	root.right.left.left.left = &node{employee: "KRISTINA"}
	printHierarchy(root, 0)

	var h hierarchical = root
	var emp1, emp2 string
	fmt.Println("YOU NEED TO MENTION TWO EMPLOYEE NAMES TO FIND THE NEAREST COMMON MANAGER")
	fmt.Print("Enter the name of 1st Employee : ")
	fmt.Scanln(&emp1)
	fmt.Print("Enter the name of 2nd Employee : ")
	fmt.Scanln(&emp2)

	//emp1, emp2 := "kristina", "wrights"     // for hard coded input

	lowestAncestor := h.lowestCommonManager(root, emp1, emp2)
	if lowestAncestor != nil {
		fmt.Printf("\nNearest Reporting Manager of %v and %v is : %v", emp1, emp2, lowestAncestor.employee)
	} else {
		fmt.Printf("\nEMPLOYEE NAMES ARE NOT VALID/PRESENT IN THE HIERARCHY !!!")
	}

}

func printHierarchy(n *node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "	"
		}
		format += "-----["
		level++

		printHierarchy(n.right, level)
		fmt.Printf(format+"%s\n", n.employee)
		printHierarchy(n.left, level)
	}
}

var emp1exist bool = false
var emp2exist bool = false

func (node *node) search(n *node, empName1, empName2 string) {

	if n == nil {
		return
	}
	if strings.ToLower(n.employee) == strings.ToLower(empName1) {
		emp1exist = true
	}
	if strings.ToLower(n.employee) == strings.ToLower(empName2) {
		emp2exist = true
	}
	node.search(n.left, empName1, empName2)
	node.search(n.right, empName1, empName2)
}

func (node *node) lowestCommonManager(n *node, emp1Name, emp2Name string) *node {
	if n == nil {
		return nil
	}

	//converting the inputs to lower case for matching
	emp1NameLower := strings.ToLower(emp1Name)
	emp2NameLower := strings.ToLower(emp2Name)

	node.search(n, emp1NameLower, emp2NameLower)
	if !emp1exist || !emp2exist {
		return nil
	}

	//condition when any of the inputs is empty string
	if emp1NameLower == "" || emp2NameLower == "" {
		fmt.Println("Please input 2 employee names")
		return nil
	}

	//condition when both inputs are same
	if emp2NameLower == emp1NameLower {
		fmt.Println("Input is ambiguous")
		return nil
	}
	if strings.ToLower(n.employee) == emp1NameLower || strings.ToLower(n.employee) == emp2NameLower {
		return n
	}
	left := node.lowestCommonManager(n.left, emp1NameLower, emp2NameLower)
	right := node.lowestCommonManager(n.right, emp1NameLower, emp2NameLower)

	if left != nil && right != nil {
		return n
	} else {
		if left != nil {
			return left
		} else {
			return right
		}

	}
}
