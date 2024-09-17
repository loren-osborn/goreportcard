//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint"
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/fzipp/gocyclo/cmd/gocyclo"
	_ "github.com/gordonklaus/ineffassign"
	_ "github.com/mgechev/revive"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
