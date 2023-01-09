# Contributor's Guide <!-- omit in toc -->

## Table of Contents <!-- omit in toc -->

- [Getting Started](#getting-started)
- [Running the CLI](#running-the-cli)
- [Testing](#testing)
- [Naming](#naming)
- [Terminology](#terminology)
- [CLI Error Handling](#cli-error-handling)
- [Testing Philosophy](#testing-philosophy)

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
go build -o create-fullstack.exe
create-fullstack.exe version

# Or:
go install . && $HOME/go/bin/create-fullstack --help
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

## Naming

All names should be `camelCase`. Do not all uppercase acronyms/abbreviations. For example:

```bash
# Do this:
HtmlRequest, ApiGenerator

# Not this:
HTMLRequest, APIGenerator
```

## Terminology

- Use API over Backend
  - This is because backend can refer to an API or CLI backend.
- Use UI over Frontend

## CLI Error Handling

<!-- TODO -->

## Testing Philosophy

- Unit tests should test implementation only. Does the function do what you expect it to produce?
  - I.e. is the function/class add a dockerfile in the way you expected?
- Integration tests should test functionality. Does the function result in the overall outcome you expect?
  - I.e. does adding a dockerfile allow you to successfully `docker build`?

Testing this way allows us to create lean unit tests, while also making room for testing functionality.
