package main

import (
	"fmt"
	"time"
)

var _ = time.Second

func main() {
	// START NO OMIT

	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "pow"
	}()

	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)

	// END NO OMIT
}
