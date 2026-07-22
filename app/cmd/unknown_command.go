package cmd

import "fmt"

type UnknownCommand struct{}

func (s UnknownCommand) Execute(args []string) {
	fmt.Printf("%s: command not found\r\n", args[0])
}
