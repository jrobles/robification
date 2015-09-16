package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postToEndpoint(url string, p []byte) string {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(p))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("-----------------------------------------------------")
	fmt.Println("RESPONSE STATUS:", resp.Status)
	fmt.Println("RESPONSE HEADERS:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("RESPONSE BODY:", string(body), string(p))
	return resp.Status
}
