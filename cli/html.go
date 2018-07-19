package cli

import (
	"os"
	"time"

	"github.com/Wraparound/wrap/ast"
	"github.com/Wraparound/wrap/html"
	"github.com/Wraparound/wrap/parser"
	"github.com/spf13/cobra"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Use:              "html [path to input file] [> output path]",
	Short:            "Export file as an HTML webpage",
	Args:             cobra.MaximumNArgs(1),
	TraverseChildren: true,
	Long:             longDescription,
	Run:              htmlRun,
}

var (
	htmlEmbedableFlag      bool
	htmlNoscenenumbersFlag bool
)

func init() {
	htmlCmd.Flags().BoolVarP(&htmlEmbedableFlag, "embedable", "e", false, "only output the play itself")
	htmlCmd.Flags().BoolVarP(&htmlNoscenenumbersFlag, "noscenenumbers", "s", false, "remove scenenumbers from output")

	WrapCmd.AddCommand(htmlCmd)
}

func htmlRun(cmd *cobra.Command, args []string) {
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
		// TODO: Make unique
		output = getOuput("script", "html")

	} else {
		pathToFile := args[0]

		if isWrapFile(pathToFile) {
			parser.UseWrapExtensions = true
		}

		script, err = parser.ParseFile(pathToFile)
		handle(err)

		// Get the file to use during export.
		output = getOuput(pathToFile, "html")
	}

	// Make sure to close the stream...
	defer output.Close()

	startExportTime := time.Now()

	if htmlNoscenenumbersFlag {
		html.AddSceneNumbers = false
	}

	if htmlEmbedableFlag {
		html.WriteHTML(script, output)

	} else {
		html.WriteHTMLPage(script, output)
	}

	endTime := time.Now()

	if benchmarkFlag {
		printBenchmarks(startTime, startExportTime, endTime)
	}
}
