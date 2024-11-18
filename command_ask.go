package main

import (
	"fmt"
	"strings"
)

func callbackAsk(cfg *config, args ...string) error {
	if cfg.Variables.Outdated {
		outdated()
		return nil
	}
	if cfg.Variables.Api == "" {
		fmt.Println("You need to be logged in!")
		return nil
	}
	if len(args) < 1 {
		fmt.Println("Please ask question")

	}
	question := strings.Join(args, " ")
	fmt.Println("Working on respond...")
	// Making request to the server
	serverResp, err := cfg.assistantClient.Ask("gpt", cfg.Variables.Api, question)
	if err != nil {
		return err
	}
	// Checking if we received valid reply from the server
	if serverResp.Reply.Valid {
		fmt.Println()
		fmt.Printf("Respond: \n%s\n", serverResp.Reply.String)
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Sorry, nothing to respond.")
		fmt.Println()

	}
	return nil
}
