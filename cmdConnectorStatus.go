package main

import (
	"log"
	"fmt"
	"github.com/spf13/cobra"
)

var cmdConnectorStatus = &cobra.Command{
	Use:   "status [name]",
	Short: "Get connector status",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println(cmd.Use)
			return
		}

		url := GetKafkaConnectUrl("connectors", args[0], "status")
		resp, err := httpClient.Get(url);
		if err != nil {
			log.Fatal(err)
		}
		HandleResponse(resp)
	},
}
