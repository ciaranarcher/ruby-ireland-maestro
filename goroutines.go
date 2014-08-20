package main

import (
	"fmt"
	"time"
)

var _ = time.Second

// START NO OMIT
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// synchronously.
	f("direct")

	// concurrently
	go f("goroutine")

	// anonymous
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("Waiting for all to finish...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done.")
}

// END NO OMIT
