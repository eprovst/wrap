package pdf

import (
	"io/ioutil"

	"github.com/Wraparound/wrap/ast"
)

// WritePDFFile writes the output of MakePDF()
func WritePDFFile(script *ast.Script, pathToFile string) error {
	// First convert file
	pdf, err := MakePDF(script)

	if err != nil {
		return err
	}

	filecontents, err := pdf.GetBytesPdfReturnErr()

	if err != nil {
		return err
	}

	return ioutil.WriteFile(pathToFile, filecontents, 0666)
}
