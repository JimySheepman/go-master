package tutorial02

import (
	"fmt"
	"io"
)

func StringifyTo[T fmt.Stringer](w io.Writer, p T) {
	fmt.Fprintln(w, p.String())
}

type Intish interface {
	~int
}

func IsPositive[T Intish](v T) bool {
	return v > 0
}

func IsGreater[T interface{ Greater(T) bool }](x, y T) bool {
	return x.Greater(y)
}
