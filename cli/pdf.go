package cli

import (
	"time"

	"github.com/Wraparound/wrap/parser"
	"github.com/Wraparound/wrap/pdf"
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
	pathToFile = getOuputPath(pathToFile)

	startExportTime := time.Now()

	if pdfNoscenenumbersFlag {
		pdf.AddSceneNumbers = false
	}

	err = pdf.MakePDF(script, pathToFile)
	handle(err)

	endTime := time.Now()

	if benchmarkFlag {
		printBenchmarks(startTime, startExportTime, endTime)
	}
}
