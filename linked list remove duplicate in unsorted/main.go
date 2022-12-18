package main

import "fmt"

type Linkedlist struct {
	head *Node
	len  int
}

type Node struct {
	val  int
	next *Node
}

func main() {
	ll := Linkedlist{}
	ll.insert(10)
	ll.insert(20)
	ll.insert(30)
	ll.insert(5)
	ll.insert(5)
	ll.insert(10)
	//ll.deleteDuplicates()
	ll.printll()
}

func (ll *Linkedlist) printll() {
	ptr := ll.head
	for i := 0; i < ll.len; i++ {
		fmt.Println(ptr.val)
		ptr = ptr.next
	}
}

// Remove duplicate nodes in an unsorted linked list?
func (ll *Linkedlist) insert(key int) {
	newNode := &Node{}
	newNode.val = key
	if ll.len == 0 {
		ll.head = newNode
		ll.len++
		return
	}
	ptr := ll.head
	for i := 0; i < ll.len; i++ {
		if ptr.next == nil {
			ptr.next = newNode
			ll.len++
			return
		}
		ptr = ptr.next
	}
}
