package main 

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type fd_new_thread struct {
	Flow_Token string `json:"flow_token"`
	Event      string `jsonevent"`
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

func fdNewThread(data fd_new_thread) (string,[]byte) {

	url := "https://i.flowdock.com/messages"

	//payload := &fdInboxStruct{Flow_Token: data.Flow_Token, Event: data.Event, External_Thread_Id: data.External_Thread_Id, Title: data.Subject}
	payload := "wait"
	p, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(p))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body),string(p))

	return resp.Status,body
}

func fdAddToThread() {
}
