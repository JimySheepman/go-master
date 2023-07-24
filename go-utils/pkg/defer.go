package main

import (
	"fmt"
	"log"
	"time"
)

func deferBehavior() error {
	start := time.Now()
	defer log.Println(time.Now().Sub(start))
	defer func() {
		log.Println("log", time.Now().Sub(start))
	}()

	// end := start.Add(12 * time.Millisecond)

	time.Sleep(1 * time.Second)

	latency := "test"
	if latency == "" {
		return fmt.Errorf("error")
	}

	fmt.Printf("latency: %s\n", latency)
	fmt.Printf("latency: %v\n", latency)
	//	fmt.Printf("latency: %d\n", latency)

	// log.Println("log", time.Now().Sub(start))
	return nil
}
