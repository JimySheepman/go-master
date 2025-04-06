package gmethods

import "fmt"

type Slice[T any] []T

func (s Slice[T]) First() T {
	return s[0]
}

func Example() {
	fmt.Println(Slice[int]([]int{1, 2, 3}).First())
	fmt.Println(Slice[string]([]string{"a", "b", "c"}).First())
}
