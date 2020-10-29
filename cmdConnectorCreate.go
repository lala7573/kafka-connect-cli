package main

import (
	"log"
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/spf13/cobra"
)


var cmdConnectorCreate = &cobra.Command{
	Use:   "create [name] [file(.json|.properties)]",
	Short: "Create connector",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 1{
			fmt.Println(cmd.Use)
			return
		}
		name, filename := args[0], args[1]
		url := GetKafkaConnectUrl("connectors")
		config, err := GetConfigFromFile(name, filename)
		if err != nil {
			log.Fatal(err)
			return
		}
		
		jsonBytes, _ := json.Marshal(&config)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			log.Fatal(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		HandleResponse(resp)
	},
}
