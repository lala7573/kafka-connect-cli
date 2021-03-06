# kafka-connect-cli

A command line interface (CLI) around the Kafka Connect REST Interface to manage connectors without any dependency.

## Installation

```
brew install lala7573/homebrew-tap/kafka-connect-cli
```

## Usage

```bash
Usage:
  ./kafka-connect-cli [command]

Available Commands:
  create      Create connector
  delete      Delete connector
  get         Get connector config
  help        Help about any command
  list        List connectors
  pause       Pause connector
  plugins     Get installed plugin list
  restart     Restart connector
  resume      Resume connector
  status      Get connector status
  update      Update connector
  validate    Validate connector config

Flags:
  -e, --endpoint string   kafka connect rest (default "http://localhost:8083/")
  -f, --format string     format: properties | json (default "properties")
  -h, --help              help for ./kafka-connect-cli

Use "./kafka-connect-cli [command] --help" for more information about a command.
```

## Build

```bash
make
# to develop
go build .
```
