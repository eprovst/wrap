//go:generate goversioninfo
// ^^ add icon to Windows build

package main

import (
	"fmt"
	"os"

	"github.com/Feltix/feltix/cmd"
)

func main() {
	// Run the root command
	if err := cmd.FeltixCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
