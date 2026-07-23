package cmd

import "os"

type ExitCommand struct{}

func (c ExitCommand) Execute(_ []string) {
	os.Exit(0)
}
