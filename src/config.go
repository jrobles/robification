package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SendGrid SendGrid `json:"sendGrid,omitempty"`
	Twilio   Twilio   `json:"twilio,omitempty"`
	Flowdock Flowdock `json:"flowdock,omitempty"`
}

type Flowdock struct {
	BotName          string `json:"botName,omitempty"`
	ChatURL          string `json:"chatURL,omitempty"`
	InboxURL         string `json:inboxURL,omntempty`
	DetailedInboxURL string `json:inboxDetailURL,omtempty`
}

type Twilio struct {
	User string `json:"accountSid,omitempty"`
	Key  string `json:"authToken,omitempty"`
}

type SendGrid struct {
	User string `json:"user,omitempty"`
	Key  string `json:"key,omitempty"`
}

func getConfig(jsonFile string) (config *Config) {
	config = &Config{}
	J, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(J), &config)
	return config
}
