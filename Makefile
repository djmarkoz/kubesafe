default: test build

clean:
	-rm -rf bin/*

build:
	go build -o bin/kubesafe

test:
	go test

update-dependencies:
	go get -u ./...

# See https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: default clean build test update-dependencies
