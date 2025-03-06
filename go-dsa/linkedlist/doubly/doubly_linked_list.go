package doubly

import "fmt"

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head == nil
}

func (dll *DoublyLinkedList) Size() int {
	size := 0
	current := dll.head

	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (dll *DoublyLinkedList) DisplayForward() {
	if dll.IsEmpty() {
		fmt.Println("Doubly linked list is empty.")
		return
	}

	current := dll.head

	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}

	fmt.Println()
}

func (dll *DoublyLinkedList) DisplayBackward() {
	if dll.IsEmpty() {
		fmt.Println("Doubly linked list is empty.")
		return
	}

	current := dll.tail

	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.prev
	}

	fmt.Println()
}

func (dll *DoublyLinkedList) InsertAtBeginning(data interface{}) {
	newNode := &Node{
		data: data,
		prev: nil,
		next: nil,
	}

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

func (dll *DoublyLinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{
		data: data,
		prev: nil,
		next: nil,
	}

	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) DeleteAtBeginning() {
	if dll.IsEmpty() {
		fmt.Println("Doubly linked list is empty. Nothing to delete.")
		return
	}

	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
}

func (dll *DoublyLinkedList) DeleteAtEnd() {
	if dll.IsEmpty() {
		fmt.Println("Doubly linked list is empty. Nothing to delete.")
		return
	}

	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}
}
