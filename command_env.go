package main

import "fmt"

func callbackEnv(cfg *config, args ...string) error {
	fmt.Println("Checking env vars")
	fmt.Println(cfg.Login)
	fmt.Println(cfg.Passwod)
	fmt.Println(cfg.Api)
	fmt.Println(cfg.UserID)

	return nil
}
