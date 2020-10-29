package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
)

var cmdConnectorResume = &cobra.Command{
	Use:   "resume [name]",
	Short: "Resume connector",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Use)
			return
		}

		url := GetKafkaConnectUrl("connectors", args[0], "resume")
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
			fmt.Println("Resumed")
		} else {
			fmt.Println(resp.StatusCode)
		}
	},
}
