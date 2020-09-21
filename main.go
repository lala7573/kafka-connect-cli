package main

import (
	"os"
	
	"github.com/spf13/cobra"
)

func getenv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
        return value
		}
	return defaultVal
}

func main() {
	cmdRoot := &cobra.Command{Use: os.Args[0]}
	cmdRoot.PersistentFlags().StringVarP(&config.KAFKA_CONNECT_REST, "endpoint", "e", getenv("KAFKA_CONNECT_REST", "http://localhost:8083/"), "kafka connect rest")
	cmdRoot.PersistentFlags().StringVarP(&config.FORMAT, "format", "f", "json", "format: properties | json")
	
	cmdRoot.AddCommand(cmdPlugins)
	cmdRoot.AddCommand(cmdConnectorList)
	cmdRoot.AddCommand(cmdConnectorGet)
	cmdRoot.AddCommand(cmdConnectorStatus)
	cmdRoot.AddCommand(cmdConnectorCreate)
	cmdRoot.AddCommand(cmdConnectorUpdate)
	cmdRoot.AddCommand(cmdConnectorDelete)
	cmdRoot.AddCommand(cmdConnectorPause)
	cmdRoot.AddCommand(cmdConnectorResume)
	cmdRoot.AddCommand(cmdConnectorRestart)
	cmdRoot.AddCommand(cmdConnectorValidate)

	// cmdRoot.AddCommand(cmdCompletion)
	cmdRoot.Execute()
}
