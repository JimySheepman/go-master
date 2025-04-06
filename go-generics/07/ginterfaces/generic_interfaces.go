package ginterfaces

import (
	"fmt"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type GenericWriter[T any] interface {
	Write(p []T) (n int, err error)
}

func Example() {
	var w Writer = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")

	var gw GenericWriter[byte] = os.Stdout
	fmt.Fprintf(gw, "hello, writer\n")
}
