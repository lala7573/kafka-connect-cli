package main

import (
	"log"
	"fmt"
	"github.com/spf13/cobra"
)

var cmdConnectorGet = &cobra.Command{
	Use:   "get [name]",
	Short: "Get connector config",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Use)
			return
		}
		url := GetKafkaConnectUrl("connectors", args[0])
		resp, err := httpClient.Get(url);
		if err != nil {
			log.Fatal(err)
			return
		}
		HandleResponse(resp)
	},
}
