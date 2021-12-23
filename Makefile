.PHONY: build
build:
	go build .
	make linux
	make mac

.PHONY: linux
linux: 
	env GOOS=linux GOARCH=amd64 go build -o kafka-connect-cli.linux .

.PHONY: darwin
mac: 
	env GOOS=darwin GOARCH=amd64 go build -o kafka-connect-cli.darwin .