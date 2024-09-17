#!/bin/sh

go install github.com/golangci/golangci-lint@latest

go install github.com/mgechev/revive@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install github.com/gordonklaus/ineffassign@latest
go install github.com/client9/misspell/cmd/misspell@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
