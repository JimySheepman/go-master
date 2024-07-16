package errorhandling

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type TimeoutError struct {
	Operation string
}

func (e TimeoutError) Error() string {
	return fmt.Sprintf("%s operation timed out", e.Operation)
}

func process(data string) error {
	timeout := 1 * time.Second
	select {
	case <-time.After(timeout):
		return TimeoutError{"Processing"}
	default:
		fmt.Println("Processing data:", data)
		return nil
	}
}

func ErrorCustom() {
	err := process("some data")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func ErrorChecking() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	// Dosyayı okuma işlemi
}

func recoverDemo() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

func PanicAndRecover() {
	defer recoverDemo()

	fmt.Println("Start processing...")
	panic("Something went wrong!")
	fmt.Println("End processing...")
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func ErrorInterface() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

func PrintErrorHandling() {
	ErrorInterface()
	ErrorCustom()
	ErrorChecking()
	PanicAndRecover()
}
