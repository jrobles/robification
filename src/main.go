package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Payload struct {
	Targets []struct {
		Destination_Type     string `json:destination_type`
		Destination_Sub_Type string `json:destination_sub_type`
		Data                 json.RawMessage
	} `json:targets`
}

type Status struct {
	Status string
}

type Responses struct {
	Messages []Status
}

var (
	config = getConfig("config.json")
)

func main() {

	// Some routing
	http.HandleFunc("/v1/ping", ping)
	http.HandleFunc("/send", sendMessage)

	// Flowdock
	http.HandleFunc("/v1/flowdock/chat", flowdockV1Chat)
	//http.HandleFunc("/v1/flowdock/inbox/basic", flowdockV1BasicInboxMessage)
	//http.HandleFunc("/v1/flowdock/inbox/detailed", flowdockV1DetaidInboxMessage)

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		fmt.Println(err)
		log.Printf("ERROR: could not start server - %v", err)
	}
}

func ping(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(res, "pong")
	log.Print("INFO: Pinged")
	res.WriteHeader(200)
}

func sendMessage(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var p Payload
	err := decoder.Decode(&p)

	statuses := []Status{}
	response := Responses{statuses}

	if err != nil {
		payloadCheck := Status{Status: "Invalid JSON"}
		response.Messages = append(response.Messages, payloadCheck)
		res.WriteHeader(400)
	} else {
		for _, payloadData := range p.Targets {
			switch string(payloadData.Destination_Type) {
			case "email":
				data := &new_email{}
				json.Unmarshal(payloadData.Data, &data)
				sendEmail(config, data)
			case "flowdock":
				switch string(payloadData.Destination_Sub_Type) {
				case "inbox_basic":
					data := &fdBasicInboxMessage{}
					json.Unmarshal(payloadData.Data, &data)
					result := Status{Status: fd_sendBasicInboxMessage(data)}
					response.Messages = append(response.Messages, result)
				case "inbox_detailed":
					data := &fdDetailedInboxMessage{}
					json.Unmarshal(payloadData.Data, &data)
					result := Status{Status: fd_sendDetailedInboxMessage(data)}
					response.Messages = append(response.Messages, result)
				case "chat":
					data := &fdChat{}
					data.External_User_Name = "robiBot"
					json.Unmarshal(payloadData.Data, &data)
					result := Status{Status: fd_sendChat(data)}
					response.Messages = append(response.Messages, result)
				}
			}
		}
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(response)
	fmt.Fprintf(res, string(b))
}
