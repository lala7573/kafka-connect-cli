package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/cobra"
)

var cmdConnectorList = &cobra.Command{
	Use:   "list",
	Short: "List connectors",
	Run: func(cmd *cobra.Command, args []string) {
		url := GetKafkaConnectUrl("connectors")
		resp, err := httpClient.Get(url);
		if err != nil {
			log.Fatal(err)
			return
		}

		body, err := ioutil.ReadAll(resp.Body);
		defer resp.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}

		var arr []string
		if err = json.Unmarshal(body, &arr); err != nil {
			log.Fatal("Failed to print json", err)
			return
		}
		if len(arr) == 0 {
			fmt.Println("no connectors.")
			return
		}

		for _, connector := range arr {
			fmt.Println(connector)
		}
	},
}

// type ConnectorResponse struct {
// 	State string `json:"state"`
// 	WorkerId string `json:"worker_id"`
// }

// type TaskResponse struct {
// 	Id int `json:"id"`
// 	State string `json:"state"`
// 	WorkerId string `json:"worker_id"`
// }

// type StatusResponse struct {
// 	Name string `json:"name"`
// 	Type string `json:"type"`
// 	Connector ConnectorResponse `json:"connector"`
// 	Tasks []TaskResponse `json:"tasks"`
// }