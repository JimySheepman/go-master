package gintro

import "fmt"

func FirstInt(s []int) int {
	return s[0]
}

func FirstString(s []string) string {
	return s[0]
}

func First[T interface{}](s []T) T {
	return s[0]
}

func Example() {
	fmt.Println(FirstInt([]int{1, 2, 3}))
	fmt.Println(FirstString([]string{"a", "b", "c"}))

	fmt.Println(First([]int{1, 2, 3}))
	fmt.Println(First([]string{"a", "b", "c"}))
}
