package main

import (
	"fmt"
	"time"
)

func main() {
	timer(time.After(2 * time.Second))
}

func timer(c <-chan time.Time) {
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}
}

/*
func timer(c <-chan time.Time, message string) {
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now(), message)
			time.Sleep(1 * time.Second)
		}
	}
}

func timer(c <-chan time.Time, message ...string) {
	defer fmt.Println("bir sornraki aşamaya geciliyor")
	defer fmt.Println("timer bitti")
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now(), message)
			time.Sleep(1 * time.Second)
		}
	}
}

func timer(c <-chan time.Time, message ...string) {
	defer fmt.Println("bir sornraki aşamaya geciliyor")
	defer fmt.Println("timer bitti")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now(), message)
			time.Sleep(1 * time.Second)

			if time.Now().Day() == 13 {
				panic("bir hata olustu..")
			}
		}
	}
}
*/
