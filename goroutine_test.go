package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

// the goroutine is a way to implement async in golang (concurency)
// to use the go routine just write "go" before the function
// this method is not suitable in function with return, because the return value will doesn't mean anything
func TestCreateGoroutines(t *testing.T) {
	//without goroutines
	// RunHelloWorld()
	// fmt.Println("Ups")
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
