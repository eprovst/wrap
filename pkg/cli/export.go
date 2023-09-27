package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/eprovst/wrap/pkg/ast"
	"github.com/eprovst/wrap/pkg/parser"
)

func export(args []string, targetExtention string, exportFunction func(*ast.Script, io.Writer) error) {
	startTime := time.Now()

	var (
		err    error
		output *os.File
		script *ast.Script
	)

	if len(args) == 0 {
		// Assume Wrap input
		parser.UseWrapExtensions = true

		script, err = getScriptFromStdin()
		handle(err)

		// Get the file to use during export.
		path, err := makeUnique("script", targetExtention)
		handle(err)

		output = getOuput(path, targetExtention)

	} else {
		pathToFile := args[0]

		if isWrapFile(pathToFile) {
			parser.UseWrapExtensions = true
		}

		script, err = parser.ParseFile(pathToFile)
		handle(err)

		// Get the file to use during export.
		output = getOuput(pathToFile, targetExtention)
	}

	// Make sure to close the stream...
	defer output.Close()

	// Make a write buffer
	buffer := bufio.NewWriter(output)

	startExportTime := time.Now()

	handle(exportFunction(script, buffer))

	handle(buffer.Flush())

	endTime := time.Now()

	if benchmarkFlag {
		printBenchmarks(startTime, startExportTime, endTime)
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
	// Get the file to use during export.
	if outFlag != "" {
		out, err := os.Create(outFlag)
		handle(err)

		return out
	}

	// No output specified, is another program expecting our output?
	if stdOutPiped() {
		return os.Stdout
	}

	// None of the above, make a new file with the same name as the source
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
	fmt.Fprintln(os.Stderr, "Parsing:  ", startExport.Sub(start))
	fmt.Fprintln(os.Stderr, "Exporting:", end.Sub(startExport))
	fmt.Fprintln(os.Stderr, "Total:    ", end.Sub(start))
}
