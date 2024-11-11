package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    callbackHelp,
		},
		"register": {
			name:        "register",
			description: "Registers new user",
			callback:    callbackRegister,
		},
		"health": {
			name:        "health",
			description: "checking on server status",
			callback:    callbackHealth,
		},
		"env": {
			name:        "env",
			description: "checking environment variables",
			callback:    callbackEnv,
		},
	}
}
