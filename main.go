package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var status string

type Payload struct {
	Source  string
	Title   string
	Content string
}

type Status struct {
	Status string
}

func main() {

	// Some routing
	http.HandleFunc("/", indexAction)
	http.HandleFunc("/send", sendAction)

	err := http.ListenAndServe(":6900", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func indexAction(res http.ResponseWriter, req *http.Request) {

}

func sendAction(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var p Payload
	err := decoder.Decode(&p)

	if err != nil {
		status = "ERROR: INVALID JSON"
	} else {
		status = "SUCCESS"
	}
	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&Status{Status: status})
	fmt.Fprintf(res, string(b))

	sendActivity(p.Title)
}
