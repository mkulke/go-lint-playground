# Go lint playground

A linter that will dissallow short-form struct initialization.

Tested with go 1.22.3 amd64

## Test

```bash
go mod tidy
go test ./internal/rules
```

## Build

```bash
go build -o no-shortform-init ./cmd/linter
```

## Run

```bash
./no-shortform-init internal/rules/testdata/src/noshorthandinit/noshorthandinit.go
```
