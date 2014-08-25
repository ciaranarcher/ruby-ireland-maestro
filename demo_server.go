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

func handleCall(res http.ResponseWriter, req *http.Request) {
	params := params{}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("error reading body ", err)
	}

	err = json.Unmarshal(data, &params)
	if err != nil {
		log.Fatal("error parsing ", err)
	}

	fmt.Println(params)
	fmt.Fprintln(res, "Hello! call ID", params.CallID, "with event", params.Event)
}

func main() {
	http.HandleFunc("/", handleCall)
	http.ListenAndServe(":1234", nil)
}
