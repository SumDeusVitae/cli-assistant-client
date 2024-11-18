package main

import "fmt"

func callbackEnv(cfg *config, args ...string) error {
	fmt.Println("Checking profile:")
	fmt.Printf("	Login: %s\n", cfg.Variables.Login)
	// Checking if we have saved api
	if cfg.Variables.Api != "" {
		fmt.Println("	Api is set")
	} else {
		fmt.Println("	No api present please Register or Login")
	}
	return nil
}
