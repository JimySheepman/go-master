package structtypes

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	Data T
	next *node[T]
	prev *node[T]
}

type list[T any] struct {
	first *node[T]
	last  *node[T]
}

func (l *list[T]) add(data T) *node[T] {
	n := node[T]{
		Data: data,
		prev: l.last,
	}

	if l.first == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	l.last.next = &n
	l.last = &n

	return &n
}

type user struct {
	name string
}

func step1() {
	var lv list[user]
	n1 := lv.add(user{"bill"})
	n2 := lv.add(user{"ale"})
	fmt.Println(n1.Data, n2.Data)

	var lp list[*user]
	n3 := lp.add(&user{"bill"})
	n4 := lp.add(&user{"ale"})
	fmt.Println(n3.Data, n4.Data)
}

func Example() {
	step1()

	exercise()
}

type stack[T any] struct {
	data []T
}

func (s *stack[T]) push(value T) {
	s.data = append(s.data, value)
}

func (s *stack[T]) pop() (T, error) {
	if len(s.data) == 0 {
		var zeroValue T
		return zeroValue, errors.New("stack is empty")
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]

	s.data = s.data[:lastIndex]

	return value, nil
}

func exercise() {
	intStack := stack[int]{}
	intStack.push(10)
	intStack.push(20)
	intStack.push(30)

	val, err := intStack.pop()
	if err == nil {
		fmt.Println("Popped:", val)
	} else {
		fmt.Println(err)
	}

	val, err = intStack.pop()
	if err == nil {
		fmt.Println("Popped:", val)
	} else {
		fmt.Println(err)
	}

	strStack := stack[string]{}
	strStack.push("Alice")
	strStack.push("Bob")

	strVal, err := strStack.pop()
	if err == nil {
		fmt.Println("Popped:", strVal)
	} else {
		fmt.Println(err)
	}

	_, err = strStack.pop()
	_, err = strStack.pop()
	if err != nil {
		fmt.Println(err)
	}
}
