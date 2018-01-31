package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version information for Wrap",
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wrap v0.1.4")
	},
}

func init() {
	WrapCmd.AddCommand(versionCmd)
}
