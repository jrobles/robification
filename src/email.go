package main

import (
	"github.com/sendgrid/sendgrid-go"
)

type new_email struct {
	From       string   `json:from`
	Subject    string   `json:subject`
	Body       string   `json:body`
	Recipients []string `json:recipients`
}

func sendEmail(config *Config, data *new_email) {
	for _, v := range data.Recipients {
		sendGrid(v, data.From, data.Subject, data.Body, config.SendGrid.User, config.SendGrid.Key)
	}

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
