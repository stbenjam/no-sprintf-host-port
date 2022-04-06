all: gosprintfhostport.so
.PHONY: lint test
clean:
	rm gosprintfhostport.so

test:
	go test ./...

lint:
	golangci-lint run ./...

gosprintfhostport.so:
	go build -buildmode=plugin plugin/gosprintfhostport.go
