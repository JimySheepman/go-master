package gvariadic

import (
	"fmt"
	"io"
	"os"
)

func First[T any](s ...T) T {
	return s[0]
}

type Slice[T any] []T

func (s Slice[T]) First() T {
	return s[0]
}

type Writer[T any] struct {
	w io.Writer
}

func (w *Writer[T]) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func Second[T io.Writer](s ...T) T {
	return s[0]
}

func Example() {
	fmt.Println(First(1, 2, 3))
	fmt.Println(First("a", "b", "c"))

	fmt.Println((Slice[int]{1, 2, 3}.First()))
	fmt.Println((Slice[string]{"a", "b", "c"}.First()))

	w := Writer[int]{w: os.Stdout}
	w.Write([]byte("Hello, world!\n"))

	ww := First(os.Stdout, os.Stderr)
	fmt.Fprintln(ww, "Hello, world!")
}
