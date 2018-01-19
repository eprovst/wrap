//go:generate goversioninfo
// ^^ add icon to Windows build

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var longDescription = `About:
Feltix is an open source command line tool that is able to
convert Fountain and Feltix files into a correctly formatted
screen- or stageplay as an HTML or a PDF.`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "feltix",
	Short: "Generate HTML and/or PDF output from Fountain files",
	Long:  longDescription,
}

var (
	outFlag       string
	benchmarkFlag bool
)

func main() {
	// Define flags used by all subcommands
	rootCmd.PersistentFlags().StringVarP(&outFlag, "out", "o", "", "specify the `file` name to be used")
	rootCmd.PersistentFlags().BoolVar(&benchmarkFlag, "benchmark", false, "measure the time spend on certain tasks")

	// Run the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
