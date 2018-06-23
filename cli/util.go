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

func printBenchmarks(start, startExport, end time.Time) {
	fmt.Fprintf(os.Stderr, "Parsing:   %d ms\n", startExport.Sub(start)/time.Millisecond)
	fmt.Fprintf(os.Stderr, "Exporting: %d ms\n", end.Sub(startExport)/time.Millisecond)
	fmt.Fprintf(os.Stderr, "Total:     %d ms\n", end.Sub(start)/time.Millisecond)
}
