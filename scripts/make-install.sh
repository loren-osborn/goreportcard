#!/bin/sh

go install github.com/alecthomas/gometalinter

go install github.com/mgechev/revive
go install github.com/fzipp/gocyclo/cmd/gocyclo
go install github.com/gordonklaus/ineffassign
go install github.com/client9/misspell/cmd/misspell
go install honnef.co/go/tools/cmd/staticcheck
