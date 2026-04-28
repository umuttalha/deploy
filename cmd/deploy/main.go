package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/umuttalha/deploy/internal/cli"
)

func main() {
	err := cli.Execute()
	if err == nil {
		return
	}
	if !errors.Is(err, cli.ErrSilent) {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	os.Exit(1)
}
