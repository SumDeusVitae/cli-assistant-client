package main

import (
	"fmt"
	"strings"
)

func callbackAsk(cfg *config, args ...string) error {
	if cfg.Variables.Api == "" {
		fmt.Println("You need to be logged in!")
		return nil
	}
	if len(args) < 1 {
		fmt.Println("Please ask question")

	}
	question := strings.Join(args, " ")
	fmt.Println("Question: \n", question)
	serverResp, err := cfg.assistantClient.Ask("gpt", cfg.Variables.Api, question)
	if err != nil {
		return err
	}
	if serverResp.Reply.Valid {
		fmt.Println()
		fmt.Printf("Reply: \n%s\n", serverResp.Reply.String)
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Sorry, nothing to respond.")
		fmt.Println()

	}
	return nil
}
