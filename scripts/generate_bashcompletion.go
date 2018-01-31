package main

import "github.com/Wraparound/wrap/cli"

func main() {
	cli.WrapCmd.GenBashCompletionFile("complete.sh")
}
