package sliceconstraints

import "fmt"

type operateFunc[T any] func(t T) T

type Slice[T any] interface {
	~[]T
}

func operate[S Slice[T], T any](slice S, fn operateFunc[T]) S {
	ret := make(S, len(slice))
	for i, v := range slice {
		ret[i] = fn(v)
	}
	return ret
}

func operate2[T any](slice []T, fn operateFunc[T]) []T {
	ret := make([]T, len(slice))
	for i, v := range slice {
		ret[i] = fn(v)
	}
	return ret
}

type Numbers []int

func DoubleUnderlying(n Numbers) Numbers {
	fn := func(n int) int {
		return 2 * n
	}

	numbers := operate2(n, fn)
	fmt.Printf("%T", numbers)
	return numbers
}

func DoubleUserDefined(n Numbers) Numbers {
	fn := func(n int) int {
		return 2 * n
	}

	numbers := operate(n, fn)
	fmt.Printf("%T", numbers)
	return numbers
}

func Example() {
	n := Numbers{10, 20, 30, 40, 50}
	fmt.Println(n)

	n = DoubleUnderlying(n)
	fmt.Println(n)

	n = DoubleUserDefined(n)
	fmt.Println(n)
}
