# Contributor's Guide

## Getting Started

Make sure you have `go 1.19`+.

Run:

```bash
go install .
```

Make sure you also set up `pre-commit`:

```bash
# For managing git hooks
pip install pre-commit

# For commit lint
npm install -g @commitlint/cli @commitlint/config-conventional

# Installs the git hooks for you, so whenever you run git commit -m "xxxx", it will run the linters automatically.
# When the linters find issues, the commit fails.
pre-commit install
```

If you want to test if `pre-commit` is working, run:

```bash
pre-commit run -a
```

You should get output like:

```bash
Trim Trailing Whitespace.................................................Passed
Fix End of Files.........................................................Passed
Check Yaml...............................................................Passed
Check for added large files..............................................Passed
golangci-lint............................................................Passed
```

## Running the CLI

```bash
# Linux
go build -o create-fullstack
./create-fullstack version

# Windows
# Not tested
go build - create-fullstack.exe
create-fullstack.exe version
```

**For autocomplete you must run:**

```bash
# Linux
./create-fullstack -install-autocomplete
```

Then restart your shell.

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
