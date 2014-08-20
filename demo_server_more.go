package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	CallID int `json:"call_id"`
	Event  string
}

type Call struct {
	ID     int
	Events chan string
	Done   chan bool
}

var calls = make(map[int]Call)

func handleCall(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	data, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(data, &request)

	if err != nil {
		log.Fatal(err)
	}

	call, exists := calls[request.CallID]
	if !exists {
		call = Call{ID: request.CallID, Events: make(chan string), Done: make(chan bool)}
		calls[request.CallID] = call
		go process(call)
		fmt.Println("Created goroutine for call", call.ID)
	} else {
		fmt.Println("Using existing call goroutine for call", call.ID)
	}

	call.Events <- request.Event
	fmt.Fprintln(w, "Hello! call ID", request.CallID, "with event", request.Event)

	// TODO block on receive from call.Done
}

func process(c Call) {
	for {
		select {
		case e := <-c.Events:
			fmt.Println("call ID", c.ID, "processing event", e)
		}
	}
}

func main() {
	http.HandleFunc("/", handleCall)
	http.ListenAndServe(":1234", nil)
}
