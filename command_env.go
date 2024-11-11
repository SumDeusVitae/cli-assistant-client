package main

import "fmt"

func callbackEnv(cfg *config, args ...string) error {
	fmt.Println("\nChecking env vars")
	fmt.Printf("Login: %s\n", cfg.Variables.Login)
	fmt.Printf("Password: %s\n", cfg.Variables.Password)
	fmt.Printf("Api: %s\n", cfg.Variables.Api)
	fmt.Printf("UserID: %s\n", cfg.Variables.UserID)
	return nil
}
