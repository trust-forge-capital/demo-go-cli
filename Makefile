APP := demo-go-cli

.PHONY: run version health echo init config build clean

run:
	go run .

version:
	go run . version

health:
	go run . health

echo:
	go run . echo hello world

init:
	go run . init

config:
	go run . config

build:
	@mkdir -p bin
	go build -o bin/$(APP) .

clean:
	rm -rf bin .demo-go-cli
