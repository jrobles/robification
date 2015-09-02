package main

import (
	"github.com/sfreiberg/gotwilio"
)

type new_sms struct {
	From       string   `json:from`
	Message    string   `json:message`
	Recipients []string `json:recipients`
}

func sendSMS(config *JSONConfigData, data *new_sms) {
	for _, v := range data.Recipients {
		//twilio(v, data.From, data.Message, config.Twilio.AccountSid, config.Twilio.AuthToken)
		print(v)
	}
}

func twilio(to string, from string, message string, accountSid string, authToken string) {
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)
	twilio.SendSMS(from, to, message, "", "")
}
