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

func handleCall(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	data, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(data, &request)

	if err != nil {
		log.Fatal("error parsing ", err)
	}

	fmt.Println(request)
	fmt.Fprintln(w, "Hello! call ID", request.CallID, "with event", request.Event)
}

func main() {
	http.HandleFunc("/", handleCall)
	http.ListenAndServe(":1234", nil)
}
