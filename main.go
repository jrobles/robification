package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Targets []struct {
		Destination_Type string `json:destination_type`
		Destination_Sub_Type string `json:destination_sub_type`
		Data struct {
			New_Email new_email
			Fd_New_Thread fd_new_thread
		} `json:data`
	} `json:targets`
}

var status string
type Status struct {
	Status string
}

func main() {

	// Some routing
	http.HandleFunc("/", indexAction)
	http.HandleFunc("/send", sendAction)

	err := http.ListenAndServe(":1337", nil)
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

fmt.Println(p)

	if err != nil {
		status = "ERROR: INVALID JSON"
	} else {

		for _,payloadData := range p.Targets {
			if (payloadData.Destination_Type) == "flowdock" {
				fmt.Println(payloadData)
			}	
		}
		status = "HEY, WE STR8"
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&Status{Status: status})
	fmt.Fprintf(res, string(b))
}
