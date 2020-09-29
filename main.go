package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Collection struct {
	Info struct {
		PostmanID   string `json:"_postman_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Schema      string `json:"schema"`
	} `json:"info"`
	Item []struct {
		Name    string `json:"name"`
		Request struct {
			Auth struct {
				Type   string `json:"type"`
				Bearer []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"bearer"`
			} `json:"auth"`
			Method string        `json:"method"`
			Header []interface{} `json:"header"`
			Body   struct {
				Mode    string `json:"mode"`
				Raw     string `json:"raw"`
				Options struct {
					Raw struct {
						Language string `json:"language"`
					} `json:"raw"`
				} `json:"options"`
			} `json:"body"`
			URL struct {
				Raw      string   `json:"raw"`
				Protocol string   `json:"protocol"`
				Host     []string `json:"host"`
				Path     []string `json:"path"`
				Query    []struct {
					Key      string `json:"key"`
					Value    string `json:"value"`
					Disabled bool   `json:"disabled"`
				} `json:"query"`
			} `json:"url"`
		} `json:"request"`
		Response []interface{} `json:"response"`
	} `json:"item"`
	ProtocolProfileBehavior struct {
	} `json:"protocolProfileBehavior"`
}

func main() {

	file, _ := ioutil.ReadFile("postmanCollection.json")
	data := Collection{}
	_ = json.Unmarshal([]byte(file), &data)
	fmt.Println(data)

}
