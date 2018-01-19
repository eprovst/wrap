package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number of Feltix",
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Feltix version v0.1.3")
	},
}

func init() {
	FeltixCmd.AddCommand(versionCmd)
}
