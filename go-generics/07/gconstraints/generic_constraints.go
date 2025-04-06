package gconstraints

import (
	"fmt"
	"io"
	"os"
)

func First[T io.Writer](s []T) T {
	return s[0]
}

type Int int

func Second[T ~int](x T) {
	fmt.Println(x)
}

func Example() {
	w := First([]io.Writer{os.Stdout, os.Stderr})
	w.Write([]byte("hello"))

	Second[Int](Int(1))
}
