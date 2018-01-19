package cmd

import (
	"github.com/spf13/cobra"
)

var longDescription = `About:
Feltix is an open source command line tool that is able to
convert Fountain and Feltix files into a correctly formatted
screen- or stageplay as an HTML or a PDF.`

// FeltixCmd represents the base command when called without any subcommands
var FeltixCmd = &cobra.Command{
	Use:   "feltix",
	Short: "Generate HTML and/or PDF output from Fountain files",
	Long:  longDescription,
}

var (
	outFlag       string
	benchmarkFlag bool
)

func init() {
	// Define flags used by all subcommands
	FeltixCmd.PersistentFlags().StringVarP(&outFlag, "out", "o", "", "specify the `file` name to be used")
	FeltixCmd.PersistentFlags().BoolVar(&benchmarkFlag, "benchmark", false, "measure the time spend on certain tasks")
}
