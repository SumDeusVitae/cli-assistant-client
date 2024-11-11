package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("\nWelcomt to CLI Assistant \nCommands:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
