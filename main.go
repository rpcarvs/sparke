/*
Copyright © 2026 Rodrigo Carvalho rpcarvs@pm.me
*/
package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/rpcarvs/sparke/cmd"
)

func main() {
	if err := fang.Execute(context.Background(), cmd.NewRootCmd()); err != nil {
		os.Exit(1)
	}
}
