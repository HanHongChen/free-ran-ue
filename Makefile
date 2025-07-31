.PHONY: build ns-up ns-down ns-ran ns-ue dc-ns-up dc-ns-down dc-ns-mran dc-ns-sran dc-ns-ue dci-ns-up dci-ns-down dci-ns-mran dci-ns-sran dci-ns-ue dci-ns-iperf-a dci-ns-iperf-b

.DEFAULT_GOAL := build

# Build the binary
build:
	GOOS=linux GOARCH=amd64 go build -o build/free-ran-ue main.go

# Basic namespace
ns-up:
	./free-ran-ue-namespace.sh up

ns-down:
	./free-ran-ue-namespace.sh down

ns-ran:
	./free-ran-ue-namespace.sh ran-ns

ns-ue:
	./free-ran-ue-namespace.sh ue-ns

# DC namespace
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

# DC Iperf namespace
dci-ns-up:
	./free-ran-ue-dc-iperf-namespace.sh up

dci-ns-down:
	./free-ran-ue-dc-iperf-namespace.sh down

dci-ns-mran:
	./free-ran-ue-dc-iperf-namespace.sh mran-ns

dci-ns-sran:
	./free-ran-ue-dc-iperf-namespace.sh sran-ns

dci-ns-ue:
	./free-ran-ue-dc-iperf-namespace.sh ue-ns

dci-ns-iperf-a:
	./free-ran-ue-dc-iperf-namespace.sh iperf-a

dci-ns-iperf-b:
	./free-ran-ue-dc-iperf-namespace.sh iperf-b