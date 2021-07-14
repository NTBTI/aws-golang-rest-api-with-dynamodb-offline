.PHONY: build clean deploy

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list todos/list/main.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create todos/create/main.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/get todos/get/main.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/update todos/update/main.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/delete todos/delete/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls offline --useDocker --dockerNetwork=serverless-go-api --noAuth --printOutput --verbose