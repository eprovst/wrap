package parser

import (
	"fmt"
	"testing"

	"github.com/Wraparound/wrap/ast"
)

const debug = false

func assertMatch(t *testing.T, script string, goal *ast.Script) {
	if debug {
		fmt.Printf("GOAL:\n%#v\n", goal)
	}

	parsed, err := ParseString(script)

	if err != nil {
		t.Error("Parsing failed.")
		return
	}

	if debug {
		fmt.Printf("\nPARSED:\n%#v\n", parsed)
	}

	if !parsed.Equals(goal) {
		t.Error("Input did not match output.")
	}
}

func scriptFromElements(elements []ast.Element) *ast.Script {
	return &ast.Script{
		Elements: elements,
	}
}

/* FOUNTAIN MODE TEST */

/* Many of the tests in this section are based on the ones by Nima Yousefi and
   John August, what follows is the license notice of their tests.

    Copyright (c) 2013 Nima Yousefi & John August

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to
    deal in the Software without restriction, including without limitation the
    rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
    sell copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
    FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
    IN THE SOFTWARE.
*/

func TestBoneyard(t *testing.T) {
	UseWrapExtensions = false

	input := `A line of action.

/* This is an inline Boneyard. */

/*
This is a
multi-line
Boneyard.
*/

This is an /* internal */ boneyard.`

	output := scriptFromElements([]ast.Element{
		ast.Action(textHandler([]string{"A line of action.", ""})),
		ast.Action(textHandler([]string{"This is an  boneyard."})),
	})

	assertMatch(t, input, output)
}
