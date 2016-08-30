package main

import (
	"encoding/json"
	"io/ioutil"
)

type JSONConfigData struct {
	SendGrid SendGrid `json:"sendGrid,omitempty"`
	Twilio   Twilio   `json:"twilio,omitempty"`
}

type Twilio struct {
	User string `json:"accountSid,omitempty"`
	Key  string `json:"authToken,omitempty"`
}

type SendGrid struct {
	User string `json:"user,omitempty"`
	Key  string `json:"key,omitempty"`
}

func getConfig(jsonFile string) (config *JSONConfigData) {
	config = &JSONConfigData{}
	J, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(J), &config)
	return config
}
