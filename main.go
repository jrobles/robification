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

type Status struct {
	Name string
}

type Responses struct {
	Items []Status
}

func main() {

	// Some routing
	http.HandleFunc("/", indexAction)
	http.HandleFunc("/send", sendAction)
	http.HandleFunc("/v1/flowdock/chat", sendChatAction)

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func sendChatAction(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	} else {
		statuses := []Status{}
		response := Responses{statuses}

		data := &fd_new_chat{}
		data.External_User_Name = "robiBot"
		data.Flow_Token = string(req.Header["Token"][0])
		data.Content = string(body)
		result := Status{Name: fdNewChat(data)}
		response.Items = append(response.Items, result)

		res.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(response)
		fmt.Fprintf(res, string(b))
	}
}

func indexAction(res http.ResponseWriter, req *http.Request) {

}

func sendAction(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var p Payload
	err := decoder.Decode(&p)

	statuses := []Status{}
	response := Responses{statuses}

	if err != nil {
		payloadCheck := Status{Name: "Invalid JSON"}
		response.Items = append(response.Items, payloadCheck)
	} else {
		for _, payloadData := range p.Targets {
			switch string(payloadData.Destination_Type) {
			case "email":
				data := &new_email{}
				json.Unmarshal(payloadData.Data, &data)
				sendEmail(getConfig("config.json"), data)
			case "flowdock":
				switch string(payloadData.Destination_Sub_Type) {
				case "inbox_basic":
					data := &fd_new_inbox_basic{}
					json.Unmarshal(payloadData.Data, &data)
					result := Status{Name: fdNewInboxBasic(data)}
					response.Items = append(response.Items, result)
				case "inbox_detailed":
					data := &fd_new_inbox_detailed{}
					json.Unmarshal(payloadData.Data, &data)
					result := Status{Name: fdNewInboxDetailed(data)}
					response.Items = append(response.Items, result)
				case "chat":
					data := &fd_new_chat{}
					data.External_User_Name = "robiBot"
					json.Unmarshal(payloadData.Data, &data)
					result := Status{Name: fdNewChat(data)}
					response.Items = append(response.Items, result)
				}
			}
		}
	}

	res.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(response)
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
