package main

import (
	"encoding/json"
)

var (
	FlowdockRobotName string = "robiBot"
	FlowdockChatUrl   string = "https://api.flowdock.com/v1/messages/chat/"
	FlowdockInboxUrl  string = "https://api.flowdock.com/v1/messages/team_inbox/"
)

type fd_new_chat struct {
	Flow_Token         string `json:"flow_token"`
	Content            string `json:"content"`
	External_User_Name string `json:"external_user_name"`
}

type fd_new_inbox_basic struct {
	Flow_Token   string `json:"flow_token"`
	Subject      string `json:"subject"`
	From_Address string `json:"from_address"`
	Source       string `json:"source"`
	Content      string `json:"content"`
}

type fd_new_inbox_detailed struct {
	Flow_Token string `json:"flow_token"`
	Event      string `json:"event"`
	Author     struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"author"`
	Title              string `json:"title"`
	External_Thread_Id string `json:"external_thread_id"`
	Thread             struct {
		Title  string `json:"title"`
		Fields []struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"fields"`
		Body         string `json:"body"`
		External_Url string `json:"external_url"`
		Status       struct {
			Color string `json:"color"`
			Value string `json:"value"`
		} `json:"status"`
	} `json:"thread"`
}

func fdNewInboxDetailed(data *fd_new_inbox_detailed) {

	p, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	postToEndpoint(FlowdockInboxUrl+data.Flow_Token, p)
}

func fdNewInboxBasic(data *fd_new_inbox_basic) {

	p, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	postToEndpoint(FlowdockInboxUrl+data.Flow_Token, p)
}

func fdNewChat(data *fd_new_chat) {
	p, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	postToEndpoint(FlowdockChatUrl+data.Flow_Token, p)
}
