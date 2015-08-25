package main

import (
	"github.com/sfreiberg/gotwilio"
)

func main() {
	accountSid := "ABC123..........ABC123"
	authToken := "ABC123..........ABC123"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+15555555555"
	to := "+15555555555"
	message := "Welcome to gotwilio!"
	twilio.SendSMS(from, to, message, "", "")
}
