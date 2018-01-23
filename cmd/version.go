package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version information for Feltix",
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Feltix v0.1.3")
	},
}

func init() {
	FeltixCmd.AddCommand(versionCmd)
}
