package main

import (
	"fmt"
	"os"
)

func runRep(cfg *config) {
	args := []string{}
	// Check if we have enough arguments (at least one argument beyond the program name)
	if len(os.Args) < 2 {
		command := getCommands()["help"]
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	// Command based on the second argument
	commandName := os.Args[1]

	command, exists := getCommands()[commandName]
	if exists {
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	} else {
		fmt.Printf("Unknown command: %s\n", commandName)
		return
	}
}
