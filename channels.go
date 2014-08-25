package main

import "fmt"

func main() {
	// START NO OMIT

	// Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	// END NO OMIT
}
