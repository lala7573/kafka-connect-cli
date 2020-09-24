package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
)

var cmdConnectorRestart = &cobra.Command{
	Use:   "restart [name]",
	Short: "Restart connector",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Use)
			return
		}
		url := GetKafkaConnectUrl("connectors", args[0], "restart")
			req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		
		if resp.StatusCode == 204 {
			fmt.Println("Restart")
		} else if resp.StatusCode == 409 {
			fmt.Println("409 Rebalance is in process")
		} else {
			fmt.Println(resp.StatusCode)
		}
	},
}
