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

func moreThanOne(booleans ...bool) bool {
	// Find first
	atLeastOneTrue := false
	for _, b := range booleans {
		if b {
			// Found the second
			if atLeastOneTrue {
				return true
			}

			// We found our first true
			atLeastOneTrue = true
		}
	}

	// No two found
	return false
}

func atLeastOne(booleans ...bool) bool {
	// Find a boolean that is true
	for _, b := range booleans {
		if b {
			return true
		}
	}

	// No true found
	return false
}
