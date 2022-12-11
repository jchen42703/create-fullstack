# Contributor's Guide

## Getting Started

Make sure you have `go 1.19`+.

Run:

```bash
go install .
```

## Testing

Regular command:

```bash
go test ./...
```

With code coverage:

```bash
go test ./... -count=1 -coverpkg=./... -coverprofile=profile.cov ./... && go tool cover -func profile.cov
```

Feel free to add the `-v` flag if you want the logs.
