package html

import (
	"io/ioutil"

	"github.com/Wraparound/wrap/ast"
)

// WriteHTMLPage writes the output of HTMLPage()
func WriteHTMLPage(script *ast.Script, pathToFile string) error {
	// First convert file
	html := ToHTMLPage(script)
	return ioutil.WriteFile(pathToFile, []byte(html), 0644)
}
