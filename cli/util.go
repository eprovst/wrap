package cli

import (
	"fmt"
	"os"
)

func handle(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+err.Error())
		os.Exit(1)
	}
}
