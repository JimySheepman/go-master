package tutorial02

import (
	"maps"
	"slices"
)

func ContainsFunc[E any](s []E, f func(E) bool) bool {
	return slices.IndexFunc(s, f) >= 0
}

func Merge[M ~map[K]V, K comparable, V any](ms ...M) M {
	result := M{}
	for _, m := range ms {
		maps.Copy[map[K]V](result, m)
	}
	return result
}
