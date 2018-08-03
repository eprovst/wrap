package cli

import (
	"bufio"
	"errors"
	"os"
	"time"

	"github.com/Wraparound/wrap/ast"
	"github.com/Wraparound/wrap/parser"
	"github.com/Wraparound/wrap/pdf"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:              "pdf [path to input file]",
	Short:            "Export file as PDF",
	Args:             cobra.MaximumNArgs(1),
	TraverseChildren: true,
	Long:             longDescription,
	Run:              pdfRun,
}

var (
	pdfNoscenenumbersFlag bool
	useCourierPrime       bool
	useCourierNew         bool
	useFreeMono           bool
)

func init() {
	pdfCmd.Flags().BoolVarP(&pdfNoscenenumbersFlag, "no-scene-numbers", "s", false, "remove scenenumbers from output")
	pdfCmd.Flags().BoolVar(&useCourierPrime, "use-courier-prime", false, "force the usage of Courier Prime")
	pdfCmd.Flags().BoolVar(&useCourierNew, "use-courier-new", false, "force the usage of Courier New")
	pdfCmd.Flags().BoolVar(&useFreeMono, "use-freemono", false, "force the usage of GNU FreeMono")

	WrapCmd.AddCommand(pdfCmd)
}

func pdfRun(cmd *cobra.Command, args []string) {
	startTime := time.Now()

	// Evaluate font selection
	if useCourierPrime && useCourierNew || useCourierPrime && useFreeMono || useCourierNew && useFreeMono {
		// The fonts are mutualy exclusive so throw an error
		handle(errors.New("tried to force multiple fonts at the same time"))
	}

	if useCourierPrime {
		pdf.SelectedFont = pdf.CourierPrime

	} else if useCourierNew {
		pdf.SelectedFont = pdf.CourierNew

	} else if useFreeMono {
		pdf.SelectedFont = pdf.FreeMono

	} else {
		// Else use automatic font selection
		pdf.SelectedFont = pdf.Auto
	}

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
		path, err := makeUnique("script", "pdf")
		handle(err)

		output = getOuput(path, "pdf")

	} else {
		pathToFile := args[0]

		if isWrapFile(pathToFile) {
			parser.UseWrapExtensions = true
		}

		script, err = parser.ParseFile(pathToFile)
		handle(err)

		// Get the file to use during export.
		output = getOuput(pathToFile, "pdf")
	}

	// Make sure to close the stream...
	defer output.Close()

	// Make a write buffer
	buffer := bufio.NewWriter(output)

	startExportTime := time.Now()

	if pdfNoscenenumbersFlag {
		pdf.AddSceneNumbers = false
	}

	handle(pdf.WritePDF(script, buffer))
	handle(buffer.Flush())

	endTime := time.Now()

	if benchmarkFlag {
		printBenchmarks(startTime, startExportTime, endTime)
	}
}
