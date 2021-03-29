package main

import (
	"os"
	"fmt"
	
	"github.com/spf13/cobra"
)

func getenv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
        return value
		}
	return defaultVal
}

const version = "0.0.4"
var cmdVersion = &cobra.Command{
	Use:                   "version",
	Short:                 "Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version, "https://github.com/lala7573/kafka-connect-cli")
	},
}

func main() {
	cmdRoot := &cobra.Command{Use: os.Args[0]}
	cmdRoot.PersistentFlags().StringVarP(&config.KAFKA_CONNECT_REST, "endpoint", "e", getenv("KAFKA_CONNECT_REST", "http://localhost:8083/"), "kafka connect rest")
	cmdRoot.PersistentFlags().StringVarP(&config.FORMAT, "format", "f", "properties", "format: properties | json")
	
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

	cmdRoot.AddCommand(cmdVersion)
	cmdRoot.AddCommand(cmdCompletion)
	cmdRoot.Execute()
}
