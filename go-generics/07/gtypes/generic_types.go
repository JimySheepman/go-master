package gtypes

import (
	"io"
	"os"
)

type Writer[T any] struct {
	w io.Writer
}

func (w *Writer[T]) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func Example() {
	w := Writer[int]{
		w: os.Stdout,
	}

	w.Write([]byte("hello"))
}
