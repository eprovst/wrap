package cli

import (
	"time"

	"github.com/Wraparound/wrap/html"
	"github.com/Wraparound/wrap/parser"
	"github.com/spf13/cobra"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Use:              "html [path to input file]",
	Short:            "Export file as an HTML webpage",
	Args:             cobra.ExactArgs(1),
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

	if htmlNoscenenumbersFlag {
		html.AddSceneNumbers = false
	}

	if htmlEmbedableFlag {
		err = html.MakeHTML(script, pathToFile)

	} else {
		err = html.MakeHTMLPage(script, pathToFile)
	}

	handle(err)

	endTime := time.Now()

	if benchmarkFlag {
		printBenchmarks(startTime, startExportTime, endTime)
	}
}
