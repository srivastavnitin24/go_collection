package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) Insert(key interface{}) {
	ptr := &Node{
		next: ll.head,
		key:  key,
	}
	if ll.head != nil {
		ll.head.prev = ptr
	}
	ll.head = ptr

	l := ll.head
	for l.next != nil {
		l = l.next
	}
	ll.tail = l
}

func (ll *LinkedList) Display() {
	list := ll.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (ll *LinkedList) Reverse() {
	curr := ll.head
	var prev *Node
	ll.tail = ll.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	ll.head = prev
	Display(ll.head)
}

func main() {
	link := LinkedList{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)
	link.Insert(78)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	link.Reverse()
	fmt.Println("\n==============================\n")

	link.findMiddleNode()
}

func (ll LinkedList) findMiddleNode() {
	ptr := ll.head
	slowptr := ptr
	fastptr := ptr
	for fastptr != nil && fastptr.next != nil {
		fastptr = fastptr.next.next
		slowptr = slowptr.next
		fmt.Println("11111111111111  ", fastptr.key, ":", slowptr.key)
	}
	fmt.Println("************* :", fastptr.key, ":", slowptr.key)
}
