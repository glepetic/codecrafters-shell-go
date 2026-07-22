package cmd

type CommandStrategy interface {
	Execute(parameters []string)
}

var commands map[string]CommandStrategy

func Init() {
	commands = make(map[string]CommandStrategy)
	commands["exit"] = ExitCommand{}
}

func SelectCommand(command string) CommandStrategy {
	commandStrategy := commands[command]
	if commandStrategy == nil {
		return UnknownCommand{}
	}
	return commandStrategy
}
