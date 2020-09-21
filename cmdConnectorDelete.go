package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
)

var cmdConnectorDelete = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete connector",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		url := GetKafkaConnectUrl("/connectors", name)
		req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 204 {
			fmt.Printf("ok")
		} else {	
			fmt.Printf("%d", resp.StatusCode)
		}
	},
}
