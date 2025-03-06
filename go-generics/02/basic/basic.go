package basic

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func printNumbers(numbers []int) {
	fmt.Print("Numbers: ")

	for _, num := range numbers {
		fmt.Print(num, " ")
	}

	fmt.Print("\n")
}

func printStrings(strings []string) {
	fmt.Print("Strings: ")

	for _, str := range strings {
		fmt.Print(str, " ")
	}

	fmt.Print("\n")
}

func printAssert(v interface{}) {
	fmt.Print("Assert: ")

	switch list := v.(type) {
	case []int:
		for _, num := range list {
			fmt.Print(num, " ")
		}

	case []string:
		for _, str := range list {
			fmt.Print(str, " ")
		}
	}

	fmt.Print("\n")
}

func printReflect(v interface{}) {
	fmt.Print("Reflect: ")

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}

	fmt.Print("\n")
}

func print[T any](slice []T) {
	fmt.Print("Generic: ")
	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func marshal[T any](value T) ([]byte, error) {
	return json.Marshal(value)
}

func Example() {
	numbers := []int{1, 2, 3}
	printNumbers(numbers)
	printAssert(numbers)
	printReflect(numbers)
	print(numbers)

	strings := []string{"A", "B", "C"}
	printStrings(strings)
	printAssert(strings)
	printReflect(strings)
	print(strings)

	var user = User{
		Name: "Alice",
		Age:  30,
	}

	jsonData, err := marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
