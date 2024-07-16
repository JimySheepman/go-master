package algorithm

import "fmt"

type structures func()

var dataStructures = []structures{
	Array,
	LinkedLists,
	Stacks,
	Queues,
	Trees,
	HashTables,
}

var dataStructuresName = map[int]string{
	0: "Array",
	1: "Linked Lists",
	2: "Stacks",
	3: "Queues",
	4: "Trees",
	5: "Hash Tables",
}

func Array() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println(arr)
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(val int) {
	newNode := &Node{data: val}
	if l.head == nil {
		l.head = newNode
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}
}

func (l *LinkedList) Print() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Println("nil")
}

func LinkedLists() {
	list := LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Print()
}

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	lastIndex := len(s.items) - 1
	val := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return val
}

func Stacks() {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func Queues() {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func (n *TreeNode) Insert(val int) {
	if val <= n.data {
		if n.left == nil {
			n.left = &TreeNode{data: val}
		} else {
			n.left.Insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &TreeNode{data: val}
		} else {
			n.right.Insert(val)
		}
	}
}

func (n *TreeNode) Search(val int) bool {
	if n == nil {
		return false
	}
	if val < n.data {
		return n.left.Search(val)
	} else if val > n.data {
		return n.right.Search(val)
	}
	return true
}

func Trees() {
	tree := &TreeNode{data: 10}
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(13)
	tree.Insert(22)
	fmt.Println(tree.Search(15)) // Output: true
	fmt.Println(tree.Search(99)) // Output: false
}
func HashTables() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m["c"])
}

func PrintDataStructures() {
	for i, ds := range dataStructures {
		fmt.Println("Data structure name:", dataStructuresName[i])
		ds()
	}
}
