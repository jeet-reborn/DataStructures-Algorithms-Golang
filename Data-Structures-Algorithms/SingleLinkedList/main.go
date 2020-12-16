/* Implement the mentioned operations of single linked list.
1) Insert a node at the beginning of the linked list.
2) Insert at the beginning of the single linked list.
3) Insert at the end of the single linked list.
4) Insert a node at given position of the single linked list.
5) Remove an element from the front of the single linked list.
6) Remove an element from the tail of the single linked list.
7) Remove an element from the given position of the linked list.
8) Search a node in the single linked list.
9) Find the sun of elements present in the single linked list.
10) Find the size of the single linked list.
*/

package main

import (
	"fmt"
)

//Node store one value, i.e, data. link is a pointer to another node.
type node struct {
	data int
	next *node
}

//LinkedList struct
type linkedList struct {
	head *node
	len  int
}

// Main function.
// Uncomment the necessary lines and test the functionalities.
func main() {

	mylist := linkedList{}   // Initialize the linked list
	node1 := &node{data: 40} // Creation of nodes.
	node2 := &node{data: 30}
	node3 := &node{data: 20}
	node4 := &node{data: 10}

	// node5 := &node{data: 50}
	// node6 := &node{data: 60}
	// node7 := &node{data: 44}

	// Insertion of nodes at the beginning.
	mylist.insert_begin(node1)
	mylist.insert_begin(node2)
	mylist.insert_begin(node3)
	mylist.insert_begin(node4)

	mylist.printList()

	/*  Insertion of nodes at the end.
	mylist.insert_end(node5)
	mylist.insert_end(node5)
	mylist.insert_end(node6)

	/* Insert at certain position
	 pos := 5
	 mylist.insert_at_postion(node7, pos)
	mylist.printList()
	*/

	/* Remove the element from the front
	mylist.remove_from_front()
	 mylist.printList()
	*/

	/* Remove the element from the tail
	mylist.remove_from_tail()
	mylist.printList()
	*/

	/* Search for an element in the linked list.
	nodeTobeFound := &node{data: 50}
	found := mylist.search_node(nodeTobeFound)
	if found {
		fmt.Println("Element is present in the linked list")
	} else {
		fmt.Println("Element NOT present in the linked list")
	}
	*/

	// count := mylist.size_of_the_linked_list()
	// fmt.Println(count)

	// sum := mylist.sum_of_elements()
	// fmt.Println(sum)

	// mylist.remove_node_at_position(3)
	// mylist.printList()
}

// insert_begin : Function to insert element at the beginning of the linked list.
func (l *linkedList) insert_begin(n *node) {
	ptr := l.head
	l.head = n
	l.head.next = ptr
	l.len++
}

//insert_end : Function to insert element at the end of the linked list.
func (l *linkedList) insert_end(n *node) {
	ptr := l.head
	n.next = nil

	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

//insert_at_postion : Function to insert element at given position of the linked list.
func (l *linkedList) insert_at_postion(n *node, pos int) {
	ptr := l.head

	if pos < 0 {
		fmt.Println("Invalid position")
		return
	}
	if pos == 0 {
		l.head = n
		l.len++
		return
	}
	if pos > l.len {
		fmt.Println("Invalid position")
		return
	}

	j := 0
	for j < pos-2 {
		j++
		ptr = ptr.next
	}
	n.next = ptr.next
	ptr.next = n
	l.len++

}

//remove_from_front : Function to remove element from the front of the linked list.
func (l *linkedList) remove_from_front() {
	ptr := l.head
	l.head = ptr.next
	l.len--
}

//remove_from_tail : Function to remove element from the back of the linked list.
func (l *linkedList) remove_from_tail() {
	ptr := l.head

	for ptr.next.next != nil {
		ptr = ptr.next
	}
	ptr.next = nil
	l.len--

}

//remove_node_at_position : Function to remove an element from a certain position from the linked list.
func (l *linkedList) remove_node_at_position(pos int) {
	ptr := l.head
	cur := ptr.next

	if pos < 0 || pos == 0 || pos > l.len {
		fmt.Println("Invalid position")
		return
	}

	if pos == 1 {
		l.head = ptr.next
		l.len--
	} else {
		j := 0
		for j < pos-2 {
			cur = cur.next
			ptr = ptr.next
			j++
		}
		ptr.next = cur.next
		cur.next = nil
		l.len--
	}
}

//search_node : Function to search an element in the single linked list.
func (l *linkedList) search_node(n *node) bool {
	ptr := l.head
	found := false

	for ptr.next != nil {
		ptr = ptr.next

		if ptr.data == n.data {
			found = true
		}
	}
	return found
}

// size_of_the_linked_list : Function to count the number of nodes present in linked list.
func (l *linkedList) size_of_the_linked_list() int {
	ptr := l.head
	count := 0
	for ptr.next != nil {
		ptr = ptr.next
		count = l.len
	}
	return count
}

// sum_of_elements : Function to add all elemets present in the linked list.
func (l *linkedList) sum_of_elements() int {
	ptr := l.head

	sum := 0
	for ptr != nil {
		sum += ptr.data
		ptr = ptr.next
	}
	return sum
}

//printList : Function to display the linked list.
func (l linkedList) printList() {
	finalList := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("%d ", finalList.data)
		finalList = finalList.next
		fmt.Printf("-> ")
	}
	fmt.Printf("\n")
}
