package main

import "github.com/Wraparound/wrap/cmd"

func main() {
	cmd.WrapCmd.GenBashCompletionFile("complete.sh")
}
