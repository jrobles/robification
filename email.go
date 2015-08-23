package main

type new_email struct {
	From string `json:from`
	Subject string `json:subject`
	Body string `json:body`
	Recipients []string `json:recipients`
}
