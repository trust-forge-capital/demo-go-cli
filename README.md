# Demo Go CLI (Cobra)

Go CLI scaffold using Cobra.

## Usage

From the project root:

```bash
go run .
```

Shows help output.

```bash
go run . version
```

Prints version string:

```
demo-go-cli version v0.1.0
```

## Makefile

```bash
make run
make version
make build
make clean
```

`make build` outputs `bin/demo-go-cli`.

## Structure

```
.
├── cmd/
│   ├── root.go
│   └── version.go
├── go.mod
└── main.go
```
