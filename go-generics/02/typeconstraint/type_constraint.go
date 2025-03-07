package typeconstraint

import "fmt"

type addOnly interface {
	string | int | int8 | int16 | int32 | int64 | float64
}

func Add[T addOnly](v1 T, v2 T) T {
	return v1 + v2
}

func index[T comparable](list []T, find T) int {
	for i, v := range list {
		if v == find {
			return i
		}
	}
	return -1
}

type person struct {
	name  string
	email string
}

func (p person) match(v person) bool {
	return p.name == v.name
}

type food struct {
	name     string
	category string
}

func (f food) match(v food) bool {
	return f.name == v.name
}

type matcher[T any] interface {
	person | food
	match(v T) bool
}

func match[T matcher[T]](list []T, find T) int {
	for i, v := range list {
		if v.match(find) {
			return i
		}
	}
	return -1
}

func Example() {
	fmt.Println(Add(10, 20))
	fmt.Println(Add("A", "B"))
	fmt.Println(Add(3.14159, 2.96))

	durations := []int{5000, 10, 40}
	findDur := 10

	i := index(durations, findDur)
	fmt.Printf("Index: %d for %d\n", i, findDur)

	people := []person{
		{
			name:  "bill",
			email: "bill@email.com",
		},
		{
			name:  "jill",
			email: "jill@email.com",
		},
		{
			name:  "tony",
			email: "tony@email.com",
		},
	}

	findPerson := person{
		name:  "tony",
		email: "tony@email.com",
	}

	i = index(people, findPerson)
	fmt.Printf("Index: %d for %s\n", i, findPerson.name)

	foods := []food{
		{
			name:     "apple",
			category: "fruit",
		},
		{
			name:     "carrot",
			category: "veg",
		},
		{
			name:     "chicken",
			category: "meat",
		},
	}
	findFood := food{
		name: "apple",
	}

	i = match(foods, findFood)
	fmt.Printf("Match: Idx: %d for %s\n", i, findFood.name)

	ints := []int{1, 2, 3, 4, 5}
	intsCopy := copyfy(ints)
	fmt.Println("Original int slice:", ints)
	fmt.Println("Copied int slice:", intsCopy)

	strs := []string{"apple", "banana", "cherry"}
	strsCopy := copyfy(strs)
	fmt.Println("Original string slice:", strs)
	fmt.Println("Copied string slice:", strsCopy)
}

type Copyable interface {
	int | string
}

func copyfy[T Copyable](input []T) []T {
	copySlice := make([]T, len(input))
	copy(copySlice, input)
	return copySlice
}
