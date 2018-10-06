package cli

import (
	"io"

	"github.com/Wraparound/wrap/ast"
	"github.com/Wraparound/wrap/html"
	"github.com/spf13/cobra"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Use:              "html [path to input file]",
	Short:            "Export file as an HTML webpage",
	Args:             cobra.MaximumNArgs(1),
	TraverseChildren: true,
	Long:             longDescription,
	Run:              htmlRun,
}

var (
	htmlEmbedableFlag  bool
	htmlProductionFlag bool
)

func init() {
	htmlCmd.Flags().BoolVarP(&htmlEmbedableFlag, "embedable", "e", false, "only output the play itself")
	htmlCmd.Flags().BoolVarP(&htmlProductionFlag, "production", "p", false, "remove scene numbers from output")

	WrapCmd.AddCommand(htmlCmd)
}

func htmlRun(cmd *cobra.Command, args []string) {
	html.Production = htmlProductionFlag

	if htmlEmbedableFlag {
		export(args, "html", func(script *ast.Script, buffer io.Writer) error {
			html.WriteHTML(script, buffer)
			return nil
		})

	} else {
		export(args, "html", func(script *ast.Script, buffer io.Writer) error {
			html.WriteHTMLPage(script, buffer)
			return nil
		})
	}
}
