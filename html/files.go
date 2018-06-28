package html

import (
	"bufio"
	"os"

	"github.com/Wraparound/wrap/ast"
)

// MakeHTML writes the output of WriteHTML() to a file
func MakeHTML(script *ast.Script, pathToFile string) error {
	// Open output file
	out, err := os.Create(pathToFile)

	if err != nil {
		return err
	}

	defer out.Close()

	buffer := bufio.NewWriter(out)
	WriteHTML(script, buffer)
	return buffer.Flush()
}

// MakeHTMLPage writes the output of WriteHTMLPage() to a file
func MakeHTMLPage(script *ast.Script, pathToFile string) error {
	// Open output file
	out, err := os.Create(pathToFile)

	if err != nil {
		return err
	}

	defer out.Close()

	buffer := bufio.NewWriter(out)
	WriteHTMLPage(script, buffer)
	return buffer.Flush()
}
