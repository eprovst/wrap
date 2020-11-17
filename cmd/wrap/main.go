//go:generate goversioninfo
// ^^ add icon to Windows build

package main

import (
	"fmt"
	"os"

	"github.com/Wraparound/wrap/pkg/cli"
)

func main() {
	// Run the root command
	if err := cli.WrapCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
