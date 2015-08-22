package main

import (
	"github.com/sendgrid/sendgrid-go"
)

func sendEmail(to string, from string, subject string, content string, config *JSONConfigData) {
	sendGrid(to, from, subject, content, config.SendGrid.User, config.SendGrid.Key)
}

func sendGrid(to string, from string, subject string, content string, user string, key string) {

	sg := sendgrid.NewSendGridClient(user, key)

	message := sendgrid.NewMail()
	message.AddTo(to)
	message.SetFrom(from)
	message.SetSubject(subject)
	message.SetText(content)

	sg.Send(message)
}
