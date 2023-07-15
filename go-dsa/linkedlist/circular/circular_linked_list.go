package circular

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type CircularLinkedList struct {
	head *Node
	tail *Node
}

func (cll *CircularLinkedList) IsEmpty() bool {
	return cll.head == nil
}

func (cll *CircularLinkedList) Size() int {
	if cll.IsEmpty() {
		return 0
	}

	size := 1
	current := cll.head.next

	for current != cll.head {
		size++
		current = current.next
	}

	return size
}

func (cll *CircularLinkedList) Display() {
	if cll.IsEmpty() {
		fmt.Println("Circular linked list is empty.")
		return
	}

	current := cll.head

	for {
		fmt.Printf("%v ", current.data)
		current = current.next
		if current == cll.head {
			break
		}
	}

	fmt.Println()
}

func (cll *CircularLinkedList) InsertAtBeginning(data interface{}) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = newNode
	} else {
		newNode.next = cll.head
		cll.tail.next = newNode
		cll.head = newNode
	}
}

func (cll *CircularLinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = newNode
	} else {
		newNode.next = cll.head
		cll.tail.next = newNode
		cll.tail = newNode
	}
}

func (cll *CircularLinkedList) DeleteAtBeginning() {
	if cll.IsEmpty() {
		fmt.Println("Circular linked list is empty. Nothing to delete.")
		return
	}

	if cll.head == cll.tail {
		cll.head = nil
		cll.tail = nil
	} else {
		cll.head = cll.head.next
		cll.tail.next = cll.head
	}
}

func (cll *CircularLinkedList) DeleteAtEnd() {
	if cll.IsEmpty() {
		fmt.Println("Circular linked list is empty. Nothing to delete.")
		return
	}

	if cll.head == cll.tail {
		cll.head = nil
		cll.tail = nil
	} else {
		current := cll.head

		for current.next != cll.tail {
			current = current.next
		}

		current.next = cll.head
		cll.tail = current
	}
}
