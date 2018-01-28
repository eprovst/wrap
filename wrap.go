//go:generate goversioninfo
// ^^ add icon to Windows build

package main

import (
	"fmt"
	"os"

	"github.com/Wraparound/wrap/cmd"
)

func main() {
	// Run the root command
	if err := cmd.WrapCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
