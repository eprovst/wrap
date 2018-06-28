package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func getOuputPath(pathToSource string) string {
	// Get the filepath to use during export.
	if outFlag != "" {
		return outFlag
	}

	extension := filepath.Ext(pathToSource)
	return strings.TrimSuffix(pathToSource, extension) + ".pdf"
}

func printBenchmarks(start, startExport, end time.Time) {
	fmt.Fprintf(os.Stderr, "Parsing:   %d ms\n", startExport.Sub(start)/time.Millisecond)
	fmt.Fprintf(os.Stderr, "Exporting: %d ms\n", end.Sub(startExport)/time.Millisecond)
	fmt.Fprintf(os.Stderr, "Total:     %d ms\n", end.Sub(start)/time.Millisecond)
}
