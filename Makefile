all: gosprintfhostport.so
.PHONY: lint

clean:
	rm gosprintfhostport.so

lint:
	golangci-lint run ./...

gosprintfhostport.so:
	go build -buildmode=plugin plugin/gosprintfhostport.go