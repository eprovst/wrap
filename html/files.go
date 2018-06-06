package html

import (
	"os"

	"github.com/Wraparound/wrap/ast"
)

// MakeHTML writes the output of WriteHTML()
func MakeHTML(script *ast.Script, pathToFile string) error {
	// Open output file
	out, err := os.Create(pathToFile)

	if err != nil {
		return err
	}

	// First convert file
	WriteHTML(script, out)
	return out.Close()
}

// MakeHTMLPage writes the output of WriteHTMLPage()
func MakeHTMLPage(script *ast.Script, pathToFile string) error {
	// Open output file
	out, err := os.Create(pathToFile)

	if err != nil {
		return err
	}

	// First convert file
	WriteHTMLPage(script, out)
	return out.Close()
}
