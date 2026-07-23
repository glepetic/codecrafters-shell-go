package cmd

import "fmt"

type CommandStrategy interface {
	Execute(parameters []string)
}

type UnknownCommand struct{}

func (c UnknownCommand) Execute(args []string) {
	fmt.Printf("%s: command not found\r\n", args[0])
}

var commands map[string]CommandStrategy

func Init() {
	commands = make(map[string]CommandStrategy)
	commands["exit"] = ExitCommand{}
	commands["echo"] = EchoCommand{}
	commands["type"] = TypeCommand{}
}

func SelectCommand(command string) CommandStrategy {
	commandStrategy := commands[command]
	if commandStrategy == nil {
		return UnknownCommand{}
	}
	return commandStrategy
}

func Exists(command string) bool {
	return commands[command] != nil
}
