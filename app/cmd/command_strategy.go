package cmd

type CommandStrategy interface {
	Execute(parameters []string)
}

func SelectCommand(command string) CommandStrategy {
	m := make(map[string]CommandStrategy)
	m["exit"] = ExitCommand{}
	return m[command]
}
