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

var calls = make(map[int]*Call)

func handleCall(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	data, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(data, &request)

	if err != nil {
		log.Fatal(err)
	}

	_, exists := calls[request.CallID]
	if !exists {
		call := &Call{ID: request.CallID, Events: make(chan string), Done: make(chan bool)}
		calls[request.CallID] = call
		process(call)
	} else {
		fmt.Println("Using existing call goroutine", request.CallID)
	}

	calls[request.CallID].Events <- request.Event

	fmt.Println(request)
	fmt.Fprintln(w, "Hello! call ID", request.CallID, "with event", request.Event)

}

func process(call *Call) {
	go func() {
		select {
		case e := <-call.Events:
			fmt.Println("call ID", call.ID, "processing event", e)
		}
	}()
}

func main() {
	http.HandleFunc("/", handleCall)
	http.ListenAndServe(":1234", nil)
}
