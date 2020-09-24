package main

import (
	"log"
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/spf13/cobra"
)

var cmdConnectorUpdate = &cobra.Command{
	Use:   "update [name] [file(.json|.properties)]",
	Short: "Update connector",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 1{
			fmt.Println(cmd.Use)
			return
		}
		
		name, filename := args[0], args[1]
		url := GetKafkaConnectUrl("connectors", name, "update")
		config, err := GetConfigFromFile(name, filename)
		if err != nil {
			log.Fatal(err)
		}
		
		jsonBytes, _ := json.Marshal(&config.Config)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		HandleResponse(resp)
	},
}
