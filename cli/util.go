package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Wraparound/wrap/ast"
	"github.com/Wraparound/wrap/parser"
)

func handle(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+err.Error())
		os.Exit(1)
	}
}

func isWrapFile(pathToFile string) bool {
	extension := strings.ToLower(filepath.Ext(pathToFile))
	return extension == ".wrap"
}

func stdInPiped() bool {
	sos, _ := os.Stdin.Stat()
	return sos.Mode()&os.ModeCharDevice == 0
}

func stdOutPiped() bool {
	sos, _ := os.Stdout.Stat()
	return sos.Mode()&os.ModeCharDevice == 0
}

func getOuput(pathToSource, extension string) *os.File {
	// Get the filepath to use during export.
	if stdOutPiped() {
		return os.Stdout
	}

	ext := filepath.Ext(pathToSource)
	path := strings.TrimSuffix(pathToSource, ext) + "." + extension
	out, err := os.Create(path)
	handle(err)

	return out
}

func getScriptFromStdin() (*ast.Script, error) {
	if stdInPiped() {
		return parser.Parser(os.Stdin)
	}

	return nil, errors.New("nothing on standard input")
}

func makeUnique(filename string, extension string) (string, error) {
	// First try the normal name
	file := filename + "." + extension
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return filename, nil
	}

	// Try to find a unique name
	for i := 1; i < 256; i++ {
		filenm := filename + "_" + strconv.Itoa(i)
		file = filenm + "." + extension

		if _, err := os.Stat(file); os.IsNotExist(err) {
			return filenm, nil
		}
	}

	// We could go further, but 256 times the same file name
	// come on...
	return "", errors.New("no unique file name possible")
}

func printBenchmarks(start, startExport, end time.Time) {
	fmt.Fprintf(os.Stderr, "Parsing:   %d ms\n", startExport.Sub(start)/time.Millisecond)
	fmt.Fprintf(os.Stderr, "Exporting: %d ms\n", end.Sub(startExport)/time.Millisecond)
	fmt.Fprintf(os.Stderr, "Total:     %d ms\n", end.Sub(start)/time.Millisecond)
}
