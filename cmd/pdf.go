package cmd

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/Wraparound/parser"
	"github.com/Wraparound/pdf"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:              "pdf [path to input file]",
	Short:            "Export file as PDF",
	Args:             cobra.ExactArgs(1),
	TraverseChildren: true,
	Long:             longDescription,
	Run:              pdfRun,
}

var pdfNoscenenumbersFlag bool

func init() {
	pdfCmd.Flags().BoolVarP(&pdfNoscenenumbersFlag, "noscenenumbers", "s", false, "remove scenenumbers from output")

	WrapCmd.AddCommand(pdfCmd)
}

func pdfRun(cmd *cobra.Command, args []string) {
	pathToFile := args[0]

	if isWrapFile(pathToFile) {
		parser.UseWrapExtensions = true
	}

	startTime := time.Now()
	script, err := parser.ParseFile(pathToFile)
	handle(err)

	// Get the filepath to use during export.
	if outFlag != "" {
		pathToFile = outFlag

	} else {
		extension := filepath.Ext(pathToFile)
		pathToFile = strings.TrimSuffix(pathToFile, extension) + ".pdf"
	}

	startExportTime := time.Now()

	if pdfNoscenenumbersFlag {
		pdf.AddSceneNumbers = false
	}

	err = pdf.WritePDFFile(script, pathToFile)
	handle(err)

	endTime := time.Now()

	if benchmarkFlag {
		printBenchmarks(startTime, startExportTime, endTime)
	}
}
