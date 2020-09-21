package main

import (
	"log"
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/spf13/cobra"
)


var cmdConnectorCreate = &cobra.Command{
	Use:   "create [name] [file(.json|.properties)]",
	Short: "Create connector",
	Run: func(cmd *cobra.Command, args []string) {
		name, filename := args[0], args[1]
		url := GetKafkaConnectUrl("/connectors")
		config, err := GetConfigFromFile(name, filename)
		if err != nil {
			log.Fatal(err)
		}
		
		jsonBytes, _ := json.Marshal(&config)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
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
