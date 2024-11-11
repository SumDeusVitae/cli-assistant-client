package main

import "fmt"

func callbackHealth(cfg *config, args ...string) error {
	message, err := cfg.assistantClient.HealthCheck()
	if err != nil {
		return err
	}
	fmt.Printf("Message: %s\n", message)
	return nil

}
