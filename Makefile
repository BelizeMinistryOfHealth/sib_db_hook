PHONY: clean build

clean:
	rm -rf ./bin ./vendor

build: clean
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/moh_api_server cmd/main.go

buildMacos: clean
	export GO111MODULE=on
	env GOOS=darwin go build -ldflags="-s -w" -o bin/moh_api_server cmd/main.go