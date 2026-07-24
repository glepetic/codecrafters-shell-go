package cmd

import (
	"fmt"
	"os/exec"
)

type TypeCommand struct{}

func (c TypeCommand) Execute(args []string) {
	if len(args) > 1 {
		for _, command := range args[1:] {
			if Exists(command) {
				fmt.Printf("%s is a shell builtin\r\n", command)
				continue
			}

			exists, path := existsGlobally(command)
			if exists {
				fmt.Printf("%s is %s\r\n", command, path)
				continue
			}

			fmt.Printf("%s: not found\r\n", command)

		}
	}
}

func existsGlobally(command string) (bool, string) {
	p, err := exec.LookPath(command)
	if err == nil {
		return true, p
	}

	return false, ""
}
