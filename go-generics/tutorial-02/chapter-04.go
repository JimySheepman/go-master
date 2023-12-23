package tutorial02

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Product[T Number](x, y T) T {
	return x * y
}

func Dupes[E comparable](s []E) bool {
	seen := map[E]bool{}
	for _, v := range s {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
