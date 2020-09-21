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
	url, err := url.Parse(config.KAFKA_CONNECT_REST)
	if err != nil {
		log.Fatal(err)
	}
	p2 := append([]string{url.Path}, paths...)
	url.Path = path.Join(p2...)
	return url.String()
}

func HandleResponse(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body);
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var obj interface{}
	if err = json.Unmarshal(body, &obj); err != nil {
		log.Fatal("Failed to print json", err)
	}
	
	if config.FORMAT == "json" {
    enc := json.NewEncoder(os.Stdout)
  	enc.SetIndent("", "  ")
		enc.Encode(obj)
	} else if (config.FORMAT == "properties") {
		fmt.Println("unsupported")
	}
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
			log.Fatal("Failed to read file", err)
		}
		config := &ConnectorConfig{}
		if err = json.Unmarshal(content, &config); err != nil {
			log.Fatal("Failed to parse ", err)
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