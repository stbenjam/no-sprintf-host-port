all: plugin
.PHONY: all plugin

lint:
	golangci-lint run ./...

plugin:
	go build -buildmode=plugin plugin/gosprintfhostport.go
