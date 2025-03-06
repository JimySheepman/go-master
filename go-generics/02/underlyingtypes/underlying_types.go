package underlyingtypes

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type vectorInt []int

func (v vectorInt) last() (int, error) {
	if len(v) == 0 {
		return 0, errors.New("empty")
	}
	return v[len(v)-1], nil
}

type vectorString []string

func (v vectorString) last() (string, error) {
	if len(v) == 0 {
		return "", errors.New("empty")
	}

	return v[len(v)-1], nil
}

func step1() {
	fmt.Print("vectorInt : ")

	vInt := vectorInt{10, -1}

	i, err := vInt.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if i < 0 {
		fmt.Print("negative integer: ")
	}

	fmt.Printf("value: %d\n", i)

	// -------------------------------------------------------------------------

	fmt.Print("vectorString : ")

	vStr := vectorString{"A", "B", string([]byte{0xff})}

	s, err := vStr.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if !utf8.ValidString(s) {
		fmt.Print("non-valid string: ")
	}

	fmt.Printf("value: %q\n", s)
}

type vectorInterface []interface{}

func (v vectorInterface) last() (interface{}, error) {
	if len(v) == 0 {
		return nil, errors.New("empty")
	}

	return v[len(v)-1], nil
}

func step2() {
	fmt.Print("vectorInterface : ")

	vItf := vectorInterface{10, "A", 20, "B", 3.14}

	itf, err := vItf.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	switch v := itf.(type) {
	case int:
		if v < 0 {
			fmt.Print("negative integer: ")
		}
	case string:
		if !utf8.ValidString(v) {
			fmt.Print("non-valid string: ")
		}
	default:
		fmt.Printf("unknown type %T: ", v)
	}

	fmt.Printf("value: %v\n", itf)
}

type vector[T any] []T

func (v vector[T]) last() (T, error) {
	var zero T

	if len(v) == 0 {
		return zero, errors.New("empty")
	}

	return v[len(v)-1], nil
}

func step3() {
	fmt.Print("vector[int] : ")

	vGenInt := vector[int]{10, -1}

	i, err := vGenInt.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if i < 0 {
		fmt.Print("negative integer: ")
	}

	fmt.Printf("value: %d\n", i)

	fmt.Print("vector[string] : ")

	vGenStr := vector[string]{"A", "B", string([]byte{0xff})}

	s, err := vGenStr.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if !utf8.ValidString(s) {
		fmt.Print("non-valid string: ")
	}

	fmt.Printf("value: %q\n", s)
}

func Example() {
	step1()
	step2()
	step3()

	exercise()
}

type keymap[T any] struct {
	data map[string]T
}

func newKeymap[T any]() *keymap[T] {
	return &keymap[T]{data: make(map[string]T)}
}

func (km *keymap[T]) set(key string, value T) {
	km.data[key] = value
}

func (km *keymap[T]) get(key string) (T, bool) {
	value, found := km.data[key]
	return value, found
}

func exercise() {
	strMap := newKeymap[string]()
	strMap.set("name", "Alice")
	strMap.set("city", "New York")

	name, found := strMap.get("name")
	if found {
		fmt.Println("Name:", name)
	} else {
		fmt.Println("Key not found")
	}

	intMap := newKeymap[int]()
	intMap.set("age", 30)
	intMap.set("year", 2025)

	age, found := intMap.get("age")
	if found {
		fmt.Println("Age:", age)
	} else {
		fmt.Println("Key not found")
	}

	_, exists := intMap.get("salary")
	fmt.Println("Salary exists:", exists)
}
