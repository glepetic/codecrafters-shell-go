package cmd

import "os"

type ExitCommand struct{}

func (s ExitCommand) Execute(_ []string) {
	os.Exit(0)
}
