package main


import (
	"os"
	"fmt"
	"log"
	"path"

	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"net/url"
	"net/http"

	"github.com/magiconair/properties"
)

func GetKafkaConnectUrl(paths ...string) string {
	restUri, err := url.Parse(config.KAFKA_CONNECT_REST)
	if err != nil {
		log.Fatal(err)
	}
	p2 := append([]string{restUri.Path}, paths...)
	restUri.Path = path.Join(p2...)
	return restUri.String()
}

func HandleResponse(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body);
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	// handle if config
	var connectorConfig ConnectorConfig;
	if err = json.Unmarshal(body, &connectorConfig); err == nil && connectorConfig.Name != "" {
		if config.FORMAT == "json" {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			enc.Encode(connectorConfig)
		} else if (config.FORMAT == "properties") {
			configStr := fmt.Sprintf("name=%s\n", connectorConfig.Name);
			for k, v := range connectorConfig.Config {
				configStr += fmt.Sprintf("%s=%s\n", k, v)
			}
			fmt.Fprintf(os.Stdout, configStr)
		}
		return;
	}

	var obj interface{}
	if err = json.Unmarshal(body, &obj); err != nil {
		log.Fatal("Failed to print json", err)
		return
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(obj)
}

type ConnectorConfig struct {
	Name      string `json:"name"`
	Config    map[string]string  `json:"config"`
}

func GetConfigFromFile(name string, filename string) (*ConnectorConfig, error) {
	extension := filepath.Ext(filename)
	if extension == ".json" {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		config := &ConnectorConfig{}
		if err = json.Unmarshal(content, &config); err != nil {
			return nil, err
		}
		
		return config, nil
	} else if extension == ".properties" {
		p := properties.MustLoadFile(filename, properties.UTF8)
		c := make(map[string]string)
		for key, value := range p.Map() {
			c[key] = value
		}

		config := &ConnectorConfig{}
		config.Name = name
		config.Config = c
		return config, nil
	} 
	return nil, fmt.Errorf("Unsupported format %s", filename)
}