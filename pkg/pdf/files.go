package pdf

import (
	"bufio"
	"io"
	"os"

	"github.com/eprovst/wrap/pkg/ast"
)

// MakePDF writes the PDF to a file
func MakePDF(script *ast.Script, pathToFile string) error {
	// Open output file
	out, err := os.Create(pathToFile)

	if err != nil {
		return err
	}

	defer out.Close()

	buffer := bufio.NewWriter(out)
	err = WritePDF(script, buffer)

	if err != nil {
		return err
	}

	return buffer.Flush()
}

// WritePDF writes the PDF to a writer
func WritePDF(script *ast.Script, writer io.Writer) error {
	// First convert file
	pdf, err := buildPDF(script)

	if err != nil {
		return err
	}

	return pdf.Write(writer)
}
