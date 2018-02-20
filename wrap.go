//go:generate goversioninfo
// ^^ add icon to Windows build

package main

import (
	"fmt"
	"os"

	"github.com/Wraparound/wrap/cli"
	"github.com/Wraparound/wrap/gui"
)

func main() {
	// Is the gui flag set? TODO: Move this into the root command?
	for _, arg := range os.Args {
		if arg == "--gui" {
			gui.DrawWindow()
			os.Exit(0)
		}
	}

	// Run the root command
	if err := cli.WrapCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
