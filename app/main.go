package main

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/cmd"
	"github.com/codecrafters-io/shell-starter-go/app/util"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	cmd.Init()
	for {
		userInput := util.ReadNext()
		args := strings.Split(userInput, " ")
		command := args[0]
		cmd.SelectCommand(command).Execute(args)
	}
}
