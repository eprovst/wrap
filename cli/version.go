package cli

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version information for Wrap",
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wrap v0.3.0 on " + strings.Title(runtime.GOOS))
	},
}

func init() {
	WrapCmd.AddCommand(versionCmd)
}
