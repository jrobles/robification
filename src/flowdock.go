package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type fdChat struct {
	Flow_Token         string `json:"flow_token,omitempty"`
	Content            string `json:"content,omitempty"`
	External_User_Name string `json:"external_user_name,omitempty"`
}

type fdBasicInboxMessage struct {
	Flow_Token   string `json:"flow_token,omitempty"`
	Subject      string `json:"subject,omitempty"`
	From_Address string `json:"from_address,omitempty"`
	Source       string `json:"source,omitempty"`
	Content      string `json:"content,omitempty"`
}

type fdDetailedInboxMessage struct {
	Flow_Token string `json:"flow_token,omitempty"`
	Event      string `json:"event,omitempty"`
	Author     struct {
		Name   string `json:"name,omitempty"`
		Avatar string `json:"avatar,omitempty"`
	} `json:"author,omitempty"`
	Title              string `json:"title,omitempty"`
	External_Thread_Id string `json:"external_thread_id,omitempty"`
	Thread             struct {
		Title  string `json:"title,omitempty"`
		Fields []struct {
			Label string `json:"label,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"fields,omitempty"`
		Body         string `json:"body,omitempty"`
		External_Url string `json:"external_url,omitempty"`
		Status       struct {
			Color string `json:"color,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"status,omitempty"`
	} `json:"thread,omitempty"`
}

func fd_sendDetailedInboxMessage(data *fdDetailedInboxMessage) string {
	p, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: could not parse data - %v", data)
	}
	return postToEndpoint(config.Flowdock.DetailedInboxURL+data.Flow_Token, p)

}

func fd_sendBasicInboxMessage(data *fdBasicInboxMessage) string {
	p, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: could not parse data - %v", data)
	}
	return postToEndpoint(config.Flowdock.InboxURL+data.Flow_Token, p)
}

func fd_sendChat(data *fdChat) string {
	p, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: could not parse data - %v", data)
	}
	return postToEndpoint(config.Flowdock.ChatURL+data.Flow_Token, p)

}

func flowdockV1Chat(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
		res.WriteHeader(400)
	} else {
		statuses := []Status{}
		response := Responses{statuses}

		data := &fdChat{}
		data.External_User_Name = config.Flowdock.BotName
		data.Flow_Token = string(req.Header["Token"][0])
		data.Content = string(body)
		result := Status{Status: fd_sendChat(data)}
		response.Messages = append(response.Messages, result)

		res.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(response)
		fmt.Fprintf(res, string(b))
		res.WriteHeader(200)
	}
}
