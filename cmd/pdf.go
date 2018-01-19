package cmd

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/Feltix/feltixparser"
	"github.com/Feltix/feltixpdf"
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

	FeltixCmd.AddCommand(pdfCmd)
}

func pdfRun(cmd *cobra.Command, args []string) {
	pathToFile := args[0]

	if isFeltixFile(pathToFile) {
		feltixparser.UseFeltixExtensions = true
	}

	startTime := time.Now()
	script, err := feltixparser.ParseFile(pathToFile)
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
		feltixpdf.AddSceneNumbers = false
	}

	err = feltixpdf.WritePDFFile(script, pathToFile)
	handle(err)

	endTime := time.Now()

	if benchmarkFlag {
		printBenchmarks(startTime, startExportTime, endTime)
	}
}
