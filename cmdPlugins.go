package main

import (
	"log"
	"github.com/spf13/cobra"
)


var cmdPlugins = &cobra.Command{
	Use:   "plugins",
	Short: "Get installed plugin list",
	Run: func(cmd *cobra.Command, args []string) {
		url := GetKafkaConnectUrl("connector-plugins")
		resp, err := httpClient.Get(url);
		if err != nil {
			log.Fatal(err)
			return
		}

		HandleResponse(resp)
	},
}

