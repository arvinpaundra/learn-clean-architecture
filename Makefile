test:
	go test -v -coverage ./...

lint: test
	golangci-lint run ./...

build: test lint
	go build -o main