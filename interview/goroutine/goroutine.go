package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

func Goroutine() {
	go say("Hello")
	go say("Concurrency")
	go say("in Go")

	time.Sleep(time.Second)
}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func Channel() {
	ch := make(chan int)

	go func() {
		ch <- 42 // Channel'a veri gönderme
	}()

	val := <-ch // Channel'dan veri okuma
	fmt.Println(val)
}

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(100)
		ch <- n
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func consumer(id int, ch <-chan int) {
	for n := range ch {
		fmt.Printf("Consumer %d received: %d\n", id, n)
	}
}

func FanInFanOut() {
	ch := make(chan int)

	go producer(ch)

	for i := 0; i < 3; i++ {
		go consumer(i, ch)
	}

	time.Sleep(3 * time.Second)
}

func generator() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(100)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return ch
}

func squareWorker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func printer(out <-chan int) {
	for n := range out {
		fmt.Println(n)
	}
}

func Pipeline() {
	printer(squareWorker(generator()))
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d starting job %d\n", id, job)
		time.Sleep(time.Second) // Simulate workload
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Example result
	}
}

func WorkerPool() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func Select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from Channel 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from Channel 2:", msg2)
		}
	}
}

func PrintGoroutine() {
	Goroutine()
	Channel()
	FanInFanOut()
	Pipeline()
	WorkerPool()
	Select()
}
