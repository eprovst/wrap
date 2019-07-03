package parser

import (
	"fmt"
	"testing"

	"github.com/Wraparound/wrap/ast"
)

const debug = true

func assertMatch(t *testing.T, script string, goal *ast.Script) {

	parsed, err := ParseString(script)

	if err != nil {
		t.Error("Parsing failed.")
		return
	}

	if !parsed.Equals(goal) {
		t.Error("Input did not match output.")

		if debug {
			fmt.Printf("GOAL:\n%#v\n", goal)
			fmt.Printf("\nPARSED:\n%#v\n", parsed)
		}
	}
}

func scriptFromElements(elements []ast.Element) *ast.Script {
	return &ast.Script{
		Elements: elements,
	}
}
