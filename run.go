package main

import (
	"fmt"
	"os"
	"strings"
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
	joinedArgs := strings.Join(os.Args[2:], " ")
	args = strings.Fields(joinedArgs)
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
