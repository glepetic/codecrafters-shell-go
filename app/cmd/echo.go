package cmd

import (
	"fmt"
	"strings"
)

type EchoCommand struct{}

func (c EchoCommand) Execute(args []string) {
	var text string
	if len(args) > 1 {
		text = strings.Join(args[1:], " ")
	} else {
		text = ""
	}
	fmt.Println(text)
}
