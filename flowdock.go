package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type activityPayload struct {
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
	} `json:thread`
}

func sendActivity(title string) {

	url := "https://api.flowdock.com/messages"
	//fmt.Println("URL:>", url)

	payload := &activityPayload{Flow_Token: "85a775399adffa1f652a3bf7b4466d77", Event: "discussion", External_Thread_Id: "1234567", Title: title}

	//var payload = []byte(`{"flow_token": "85a775399adffa1f652a3bf7b4466d77","event":"message","author": {"name": "Marty","avatar": "https://avatars.githubusercontent.com/u/3017123?v=3"},"title": "updated ticket","external_thread_id": "1234567","thread": {"title": "Polish the flux capacitor","fields": [{ "label": "Dustiness", "value": "5 - severe" }],"body": "The flux capacitor has been in storage for more than 30 years and it needs to be spick and span for the re-launch.","external_url": "https://example.com/projects/bttf/tickets/1234567","status": {"color": "green","value": "open"}}}`)
	p, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(p))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(p))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
