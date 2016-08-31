package main

import (
	"encoding/json"
	"log"
)

var (
	FlowdockRobotName      string = "robiBot"
	FlowdockChatUrl        string = "https://api.flowdock.com/v1/messages/chat/"
	FlowdockInboxUrl       string = "https://api.flowdock.com/v1/messages/team_inbox/"
	FlowdockInboxDetailUrl string = "https://api.flowdock.com/v1/flows/sandboxstudio/a51/messages/"
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
	return postToEndpoint(FlowdockInboxDetailUrl+data.Flow_Token, p)

}

func fd_sendBasicInboxMessage(data *fdBasicInboxMessage) string {
	p, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: could not parse data - %v", data)
	}
	return postToEndpoint(FlowdockInboxUrl+data.Flow_Token, p)
}

func fd_sendChat(data *fdChat) string {
	p, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: could not parse data - %v", data)
	}
	return postToEndpoint(FlowdockChatUrl+data.Flow_Token, p)

}
