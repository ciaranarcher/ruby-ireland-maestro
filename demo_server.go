package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type params struct {
	CallID int `json:"call_id"`
	Event  string
}

func handleCall(w http.ResponseWriter, req *http.Request) {
	params := params{}
	data, _ := ioutil.ReadAll(req.Body)

	err := json.Unmarshal(data, &params)

	if err != nil {
		log.Fatal("error parsing ", err)
	}

	fmt.Println(params)
	fmt.Fprintln(w, "Hello! call ID", params.CallID, "with event", params.Event)
}

func main() {
	http.HandleFunc("/", handleCall)
	http.ListenAndServe(":1234", nil)
}
