package main

import "github.com/Feltix/feltix/cmd"

func main() {
	cmd.FeltixCmd.GenBashCompletionFile("complete.sh")
}
