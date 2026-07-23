package cmd

import "fmt"

type TypeCommand struct{}

func (c TypeCommand) Execute(args []string) {
	if len(args) > 1 {
		for _, command := range args[1:] {
			if Exists(command) {
				fmt.Printf("%s is a shell builtin\r\n", command)
			} else {
				fmt.Printf("%s: not found\r\n", command)
			}
		}
	}
}
