package main

import (
	"fmt"
)

func callbackWhoami(cfg *config, args ...string) error {
	if cfg.Variables.Outdated {
		outdated()
		return nil
	}

	if cfg.Variables.Api == "" {
		fmt.Println("You are not logged in please register or login")
		callbackHelp(cfg, args...)
		return nil
	}
	serverResp, err := cfg.assistantClient.Whoami(cfg.Variables.Api)
	if err != nil {
		if err.Error() == "unauthorized" {
			fmt.Println("Seems like something wrong with credentials")
			fmt.Println("Please log in or register")
			fmt.Println("Use <help> command for assistance")
			return nil
		}
		return err
	}
	fmt.Printf("You are logged as: %s\n", serverResp.Login)
	if serverResp.Email.Valid {
		fmt.Printf("Email: %s\n", serverResp.Email.String)
	}

	return nil

}
