package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JSONConfigData struct {
	SendGrid struct {
		User string `json:user`
		Key  string `json:key`
	} `json:sendGrid`
}

type Payload struct {
	Targets []struct {
		Destination_Type     string `json:destination_type`
		Destination_Sub_Type string `json:destination_sub_type`
		Data                 json.RawMessage
	} `json:targets`
}

var status string

type Status struct {
	Status string
}

func main() {

	// Some routing
	http.HandleFunc("/", indexAction)
	http.HandleFunc("/send", sendAction)

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func indexAction(res http.ResponseWriter, req *http.Request) {

}

func sendAction(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var p Payload
	err := decoder.Decode(&p)

	if err != nil {
		status = "ERROR: INVALID JSON"
	} else {

		for _, payloadData := range p.Targets {
			switch string(payloadData.Destination_Type) {
			case "email":
				data := &new_email{}
				json.Unmarshal(payloadData.Data, &data)
				sendEmail(getConfig("config.json"), data)
			case "flowdock":
				data := &fd_new_thread{}
				json.Unmarshal(payloadData.Data, &data)
				fdNewThread(data)
			case "sms":
				//data := &new_sms{}
				//json.Unmarshal(payloadData.Data, &data)
				//sendSms(getConfig("config.json"), data)
			}
		}
		status = "HEY, WE STR8"
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&Status{Status: status})
	fmt.Fprintf(res, string(b))
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
