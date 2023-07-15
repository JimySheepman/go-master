package singly

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func (sll *SinglyLinkedList) IsEmpty() bool {
	return sll.head == nil
}

func (sll *SinglyLinkedList) Size() int {
	size := 0
	current := sll.head

	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (sll *SinglyLinkedList) Display() {
	if sll.IsEmpty() {
		fmt.Println("Singly linked list is empty.")
		return
	}

	current := sll.head

	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}

	fmt.Println()
}

func (sll *SinglyLinkedList) InsertAtBeginning(data interface{}) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	newNode.next = sll.head
	sll.head = newNode
}

func (sll *SinglyLinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if sll.IsEmpty() {
		sll.head = newNode
	} else {
		current := sll.head

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (sll *SinglyLinkedList) DeleteAtBeginning() {
	if sll.IsEmpty() {
		fmt.Println("Singly linked list is empty. Nothing to delete.")
		return
	}

	sll.head = sll.head.next
}

func (sll *SinglyLinkedList) DeleteAtEnd() {
	if sll.IsEmpty() {
		fmt.Println("Singly linked list is empty. Nothing to delete.")
		return
	}

	if sll.head.next == nil {
		sll.head = nil
		return
	}

	current := sll.head

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
}
