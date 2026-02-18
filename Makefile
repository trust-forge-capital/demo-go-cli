APP := demo-go-cli

.PHONY: run version build clean

run:
	go run .

version:
	go run . version

build:
	@mkdir -p bin
	go build -o bin/$(APP) .

clean:
	rm -rf bin
