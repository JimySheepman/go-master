package tutorial02

import (
	"fmt"
	"io"
)

func PrintAnythingTo[T any](w io.Writer, p T) {
	fmt.Fprintln(w, p)
}

type Group[E any] []E

func Len[E any](s []E) int {
	return len(s)
}
