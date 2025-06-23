.PHONY: build

build:
	GOOS=linux GOARCH=amd64 go build -o build/free-ran-ue main.go