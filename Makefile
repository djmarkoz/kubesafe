default: test build

test:
	go test

build:
	go build -o bin/kubesafe

