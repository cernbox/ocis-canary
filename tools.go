// +build tools

package main

import (
	_ "github.com/UnnoTed/fileb0x"
	_ "github.com/cespare/reflex"
	_ "github.com/haya14busa/goverage"
	_ "github.com/mitchellh/gox"
	_ "github.com/restic/calens"
	_ "golang.org/x/lint/golint"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
