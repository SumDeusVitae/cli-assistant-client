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
		"login": {
			name:        "login",
			description: "login to your account",
			callback:    callbackLogin,
		},
		"q": {
			name:        "q {question}",
			description: "ask AI question, get reply",
			callback:    callbackAsk,
		},
		"whoami": {
			name:        "whoami",
			description: "checks if you are logged in",
			callback:    callbackWhoami,
		},
		"version": {
			name:        "version",
			description: "checks on version",
			callback:    callbackVer,
		},
		"update": {
			name:        "update",
			description: "updates to the latest version",
			callback:    callbackUpdate,
		},
	}
}
