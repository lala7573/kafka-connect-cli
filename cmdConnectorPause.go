package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
)

var cmdConnectorPause = &cobra.Command{
	Use:   "pause [name]",
	Short: "Pause connector",
	Run: func(cmd *cobra.Command, args []string) {
		url := GetKafkaConnectUrl("connectors", args[0], "pause")
			req, err := http.NewRequest("PUT", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 202 {
			fmt.Println("Paused")
		} else {
			fmt.Println(resp.StatusCode)
		}
	},
}
