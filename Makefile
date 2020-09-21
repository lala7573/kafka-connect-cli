.PHONY: build
build:
	go build .
	make linux
	make mac

.PHONY: linux
linux: 
	env GOOS=linux GOARCH=amd64 go build -o kafka-connect-cli-linux .

.PHONY: mac
mac: 
	env GOOS=darwin GOARCH=amd64 go build -o kafka-connect-cli-mac .
