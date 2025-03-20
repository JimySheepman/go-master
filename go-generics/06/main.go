package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer | constraints.Complex
}

func Double[T Number](value T) T {
	return value * 2
}

func DotProduct[T Number](s1, s2 []T) T {
	if len(s1) != len(s2) {
		panic("error")
	}
	var r T
	for i := range s1 {
		r += s1[i] + s2[i]
	}
	return r
}

func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func example01() {
	fmt.Println(Double(23))
	fmt.Println(Double(23.23))
	fmt.Println(Double(-2323.3434))

	i := []int{1, 2, 3}
	j := []int{4, 5, 6}
	fmt.Println(DotProduct(i, j))

	ints := map[string]int64{
		"first":   23,
		"second":  565,
		"third":   755,
		"fourth":  766,
		"fifth":   8977,
		"sixth":   70433,
		"seventh": 4339222,
	}
	fmt.Println(Sum(ints))
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Merge[T any](slices ...[]T) (mergedSlice []T) {
	for _, slice := range slices {
		mergedSlice = append(mergedSlice, slice...)
	}
	return mergedSlice
}

func Includes[T comparable](slice []T, value T) bool {
	for _, el := range slice {
		if el == value {
			return true
		}
	}
	return false
}

func Sort[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func example02() {
	s := []int{1, 2, 3, 7, 5, 22, 18}
	j := []int{4, 5, 6}

	floats := Map(s, func(i int) float64 { return float64(i) })
	fmt.Println(floats)

	sum := Reduce(s, 0, func(i, j int) int { return i + j })
	fmt.Println(sum)

	evens := Filter(s, func(i int) bool { return i%2 == 0 })
	fmt.Println(evens)

	merged := Merge(s, j)
	fmt.Println(merged)

	i := Includes(s, 22)
	fmt.Println(i)

	Sort(s)
	fmt.Println(s)
}

func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func example03() {
	k := Keys(map[int]int{1: 2, 2: 4})
	fmt.Println(k)

	s := Sum(map[int]int{1: 2, 2: 4})
	fmt.Println(s)

}

type Set[T comparable] map[T]struct{}

func Make[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}

func example04() {
	set := Make[int]()

	set.Add(1)
	set.Add(3)
	set.Add(5)
	set.Add(7)
	set.Add(1)
	fmt.Println(set.Contains(2))
}

func main() {
	example01()
	example02()
	example03()
	example04()
}
