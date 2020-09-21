package main

import (
	"os"
	"log"
	"bytes"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/cobra"
)

var cmdConnectorValidate = &cobra.Command{
	Use:   "validate [name] [file(.json|.properties)]",
	Short: "Validate connector config",
	Run: func(cmd *cobra.Command, args []string) {
		name, filename := args[0], args[1]
		config, err := GetConfigFromFile(name, filename)
		if err != nil {
			log.Fatal(err)
		}

		config.Config["name"] = name
		jsonBytes, _ := json.Marshal(&config.Config)
		className := config.Config["connector.class"]
		classNameSplit := strings.Split(className, ".")
		classNameOnly := classNameSplit[len(classNameSplit)-1]
		url := GetKafkaConnectUrl("connector-plugins", classNameOnly, "config/validate")
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body);
		defer resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		var response ValidateResponse
		if err = json.Unmarshal(body, &response); err != nil {
			log.Fatal("Failed to print json", err)
		}

		for _, config:= range response.Configs {
			if len(config.Value.Errors) > 0 {
				enc := json.NewEncoder(os.Stdout)
				enc.SetIndent("", "  ")
				enc.Encode(config.Value)
			}
		}
	},
}

type ValidateResponse struct {
	Name string `json:"name"`
	ErrorCount int `json:"error_count"`
	Groups []string `json:"groups"`
	Configs []ValidateConfigResponse `json:"configs"`
}

type ValidateConfigResponse struct {
	 Definition ValidateDefinitionResponse `json:"definition"`
	 Value ValidateValueResponse `json:"value"`
}

type ValidateDefinitionResponse struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Required bool `json:"required"`
	DefaultValue string `json:"default_value"`
	Importance string `json:"importance"`
	Documentation string `json:"documentation"`
	Group string `json:"group"`
	Width string `json:"width"`
	DisplayName string `json:"display_name"`
	Dependents []string `json:"dependents"`
	Order int `json:"order"`
}

type ValidateValueResponse struct {
	Name string `json:"name"`
	Value string `json:"value"`
	RecommendedValues []string `json:"recommended_values"`
	Errors []string `json:"errors"`
	Visible bool `json:"visible"`
}