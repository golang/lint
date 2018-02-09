# Contributing to Golint

## Before filing an issue:

### Are you having trouble building golint?

Check you have the latest version of its dependencies. Run
```
go get -u github.com/golang/lint
```
If you still have problems, consider searching for existing issues before filing a new issue.

## Before sending a pull request:

Have you understood the purpose of golint? Make sure to carefully read `README`.

### Run golint tests. Note that golint will report warnings on test files.
go test -v -race ./...
