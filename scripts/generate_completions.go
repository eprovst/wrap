package main

import "github.com/eprovst/wrap/pkg/cli"

func main() {
	cli.WrapCmd.GenBashCompletionFile("bash-complete.sh")
	cli.WrapCmd.GenZshCompletionFile("zsh-complete.sh")
}
