package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

type Node struct {
	val  int
	next *Node
}

func main() {
	ll := LinkedList{}
	ll.insert(31)
	ll.insert(22)
	ll.insert(16)
	ll.insert(100)
	ll.print()

	ll.search(16)
	ll.search(15)

	fmt.Println("--------getNodeAt------------- : ", ll.getNodeAt(2).val)
	fmt.Println("-----------getNodeAt---------- : ", ll.getNodeAt(3).val)

	ll.delete(31)
	fmt.Println("-----------length after delete ll ----------", ll.len)
	ll.print()

	fmt.Println("-----------insert at ----------")
	ll.insertAt(1, 7)
	ll.print()

	fmt.Println("\n-----------delete at ----------")
	ll.deleteAt(4)
	ll.print()

}

func (ll *LinkedList) print() {
	if ll.len == 0 {
		fmt.Println("Linked list doesn't have any node, its empty !")
	}

	ptr := ll.head
	for i := 0; i < ll.len; i++ {
		fmt.Print(" ", ptr.val)
		ptr = ptr.next
	}
}

func (ll *LinkedList) insert(key int) {
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

func (ll *LinkedList) insertAt(pos, key int) {
	newNode := &Node{}
	newNode.val = key

	if pos < 0 || pos > ll.len {
		return
	}
	if pos == 0 {
		ll.head = newNode
		ll.len++
		return
	}
	n := ll.getNodeAt(pos)
	newNode.next = n
	prev := ll.getNodeAt(pos - 1)
	prev.next = newNode
	ll.len++
	return
}

func (ll LinkedList) getNodeAt(position int) *Node {
	ptr := ll.head
	if position < 0 {
		return ptr
	}
	if position > (ll.len - 1) {
		return nil
	}
	for i := 0; i < position; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (ll *LinkedList) delete(key int) {
	ptr := ll.head
	if ll.len == 0 {
		fmt.Println("Linklist is empty cannot delete")
		return
	}
	for i := 0; i < ll.len; i++ {
		if ptr.val == key {
			if i > 0 {
				prevNode := ll.getNodeAt(i - 1)
				prevNode.next = ptr.next
				fmt.Println("\nDelete successfully")
			} else {
				ll.head = ptr.next
			}
			ll.len--
			return
		}
		ptr = ptr.next
	}
	return
}

func (ll *LinkedList) deleteAt(pos int) {

	if pos > ll.len-1 || pos < 0 {
		return
	}
	if pos == ll.len-1 {
		n := ll.getNodeAt(pos - 1)
		n.next = nil
		ll.len--
		return
	}
	n := ll.getNodeAt(pos)
	prevnode := ll.getNodeAt(pos - 1)
	prevnode.next = n.next
	ll.len--
	return

}

func (ll *LinkedList) search(key int) {
	if ll.len != 0 {
		ptr := ll.head
		for i := 0; i < ll.len; i++ {
			if ptr.val == key {
				fmt.Println("\nSearched key")
				return
			}
			ptr = ptr.next
		}
	} else {
		fmt.Println("Linklist is empty cannot search")
	}
}

//find middle node in one iteration
func (ll LinkedList) findMiddleNode() {
	ptr := ll.head
	slowptr := ptr
	fastptr := ptr
	for fastptr != nil && fastptr.next != nil {
		fastptr = fastptr.next.next
		slowptr = slowptr.next
		fmt.Println("11111111111111  ", fastptr.val, ":", slowptr.val)
	}
	fmt.Println("************* :", fastptr.val, ":", slowptr.val)
}
