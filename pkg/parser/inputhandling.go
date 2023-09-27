package parser

import (
	"bytes"
	"os"

	"github.com/eprovst/wrap/pkg/ast"
)

// ParseFile parses directly from a file path.
func ParseFile(filename string) (*ast.Script, error) {
	file, err := os.Open(filename)
	defer file.Close() // Make sure it's closed at one point.

	if err != nil {
		return nil, err
	}

	return Parser(file)
}

// ParseString parses a string.
func ParseString(script string) (*ast.Script, error) {
	// We basicaly turn the string into a file in memory.
	var buffer bytes.Buffer
	buffer.WriteString(script)

	return Parser(&buffer)
}

// ParseText parses a list of strings.
func ParseText(text []string) (*ast.Script, error) {
	var buffer bytes.Buffer

	/* Similar to ParseString but we first combine all the lines.
	   Inserting \n between them */
	for i := 0; i < len(text); i++ {
		buffer.WriteString(text[i])

		if i+1 != len(text) {
			buffer.WriteString("\n")
		}
	}

	return Parser(&buffer)
}
