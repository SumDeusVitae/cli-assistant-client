package main

import "fmt"

func callbackEnv(cfg *config, args ...string) error {
	fmt.Println("Checking profile:")
	fmt.Printf("	Login: %s\n", cfg.Variables.Login)
	// fmt.Printf("Password: %s\n", cfg.Variables.Password)
	if cfg.Variables.Api != "" {
		fmt.Println("	Api is set")
	} else {
		fmt.Println("	No api present please Register or Login")
	}
	// fmt.Printf("Api: %s\n", cfg.Variables.Api)
	// fmt.Printf("UserID: %s\n", cfg.Variables.UserID)
	return nil
}
