package main

import (
	"context"
	"fmt"
	"os"

	"github.com/r9s-ai/onr-sdk/golang/internal/cli"
)

func main() {
	if err := cli.ExecuteContext(context.Background(), os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
