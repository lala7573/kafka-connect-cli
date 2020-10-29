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
		if len(args) == 0 {
			fmt.Println(cmd.Use)
			return
		}
		url := GetKafkaConnectUrl("connectors", args[0], "pause")
			req, err := http.NewRequest("PUT", url, nil)
		if err != nil {
			log.Fatal(err)
			return
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		if resp.StatusCode == 202 {
			fmt.Println("Paused")
		} else {
			fmt.Println(resp.StatusCode)
		}
	},
}
