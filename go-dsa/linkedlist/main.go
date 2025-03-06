package main

import (
	"linked-list/circular"
	"linked-list/doubly"
	"linked-list/singly"
)

func main() {
	sll := &singly.SinglyLinkedList{}

	sll.InsertAtBeginning(3)
	sll.InsertAtBeginning(2)
	sll.InsertAtBeginning(1)

	sll.InsertAtEnd(4)
	sll.InsertAtEnd(5)
	sll.InsertAtEnd(6)

	sll.Display()

	sll.DeleteAtBeginning()
	sll.DeleteAtBeginning()

	sll.Display()

	sll.DeleteAtEnd()
	sll.DeleteAtEnd()

	sll.Display()

	cll := &circular.CircularLinkedList{}

	cll.InsertAtBeginning(3)
	cll.InsertAtBeginning(2)
	cll.InsertAtBeginning(1)

	cll.InsertAtEnd(4)
	cll.InsertAtEnd(5)
	cll.InsertAtEnd(6)

	cll.Display()

	cll.DeleteAtBeginning()
	cll.DeleteAtBeginning()

	cll.Display()

	cll.DeleteAtEnd()
	cll.DeleteAtEnd()

	cll.Display()

	dll := &doubly.DoublyLinkedList{}

	dll.InsertAtBeginning(3)
	dll.InsertAtBeginning(2)
	dll.InsertAtBeginning(1)

	dll.InsertAtEnd(4)
	dll.InsertAtEnd(5)
	dll.InsertAtEnd(6)

	dll.DisplayForward()
	dll.DisplayBackward()

	dll.DeleteAtBeginning()
	dll.DeleteAtBeginning()

	dll.DisplayForward()
	dll.DisplayBackward()

	dll.DeleteAtEnd()
	dll.DeleteAtEnd()

	dll.DisplayForward()
	dll.DisplayBackward()
}
