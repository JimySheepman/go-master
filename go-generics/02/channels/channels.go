package channels

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type doworkFn[Result any] func(context.Context) Result

func doWork[Result any](ctx context.Context, work doworkFn[Result]) chan Result {
	ch := make(chan Result, 1)

	go func() {
		ch <- work(ctx)
		fmt.Println("doWork : work complete")
	}()

	return ch
}

func Example() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	dwf := func(ctx context.Context) string {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		return "work complete"
	}

	select {
	case v := <-doWork(ctx, dwf):
		fmt.Println("main:", v)
	case <-ctx.Done():
		fmt.Println("main: timeout")
	}
}
