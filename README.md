# Demo Go CLI (Cobra)

Go CLI scaffold using Cobra.

## Features

- `version`: print CLI version
- `health`: output basic health status
- `echo`: echo input text (`--upper` supported)
- `init`: initialize local config file
- `config`: show config file status/path

## Usage

From the project root:

```bash
go run .
```

Shows help output.

```bash
go run . version
# demo-go-cli version v0.2.0

go run . health
# status=ok app=demo-go-cli version=v0.2.0 time=...

go run . echo hello team
# hello team

go run . echo --upper hello team
# HELLO TEAM

go run . init
# initialized config: .demo-go-cli/config.yaml

go run . config
# config exists: .demo-go-cli/config.yaml
```

## Makefile

```bash
make run
make version
make health
make echo
make init
make config
make build
make clean
```

`make build` outputs `bin/demo-go-cli`.

## Structure

```
.
├── cmd/
│   ├── config.go
│   ├── echo.go
│   ├── health.go
│   ├── init.go
│   ├── root.go
│   └── version.go
├── go.mod
└── main.go
```
