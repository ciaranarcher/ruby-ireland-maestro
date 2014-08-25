package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type params struct {
	CallID int `json:"call_id"`
	Event  string
}

type callProcessor struct {
	ID     int
	Events chan string
	Done   chan bool
}

var calls = make(map[int]callProcessor)

func handleCall(w http.ResponseWriter, r *http.Request) {
	params := params{}
	data, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(data, &params)

	if err != nil {
		log.Fatal(err)
	}

	call, exists := calls[params.CallID]
	if !exists {
		call = callProcessor{ID: params.CallID, Events: make(chan string), Done: make(chan bool)}
		calls[params.CallID] = call
		go process(call)
		fmt.Println("Created goroutine for call", call.ID)
	} else {
		fmt.Println("Using existing goroutine for call", call.ID)
	}

	call.Events <- params.Event
	<-call.Done
	fmt.Fprintln(w, "Request complete; call ID", params.CallID, "with event", params.Event, "processed successfully.")
}

func process(c callProcessor) {
	timeout := time.After(30 * time.Second)
	for {
		select {
		case e := <-c.Events:
			time.Sleep(1 * time.Second)
			fmt.Println("Call ID", c.ID, "processed event", e, "successfully.")
			c.Done <- true
		case <-timeout:
			delete(calls, c.ID)
			fmt.Println("Cleaning up goroutine for call ID", c.ID)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handleCall)
	http.ListenAndServe(":1234", nil)
}
