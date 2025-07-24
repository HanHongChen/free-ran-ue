.PHONY: build ns-up ns-down ns-ran ns-ue dc-ns-up dc-ns-down dc-ns-mran dc-ns-sran dc-ns-ue

.DEFAULT_GOAL := build

build:
	GOOS=linux GOARCH=amd64 go build -o build/free-ran-ue main.go

ns-up:
	./free-ran-ue-namespace.sh up

ns-down:
	./free-ran-ue-namespace.sh down

ns-ran:
	./free-ran-ue-namespace.sh ran-ns

ns-ue:
	./free-ran-ue-namespace.sh ue-ns

dc-ns-up:
	./free-ran-ue-dc-namespace.sh up

dc-ns-down:
	./free-ran-ue-dc-namespace.sh down

dc-ns-mran:
	./free-ran-ue-dc-namespace.sh mran-ns

dc-ns-sran:
	./free-ran-ue-dc-namespace.sh sran-ns

dc-ns-ue:
	./free-ran-ue-dc-namespace.sh ue-ns