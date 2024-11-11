package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type RegistrationRespond struct {
	UserID string `json:"user_id"`
	APIKey string `json:"api_key"`
}

func callbackRegister(cfg *config, args ...string) error {
	// Login and Password
	cfg.Login = enterValid("Login")
	err := os.Setenv("LOGIN", cfg.Login)
	if err != nil {
		log.Println("Couldn't save login to the environment")
	}
	log.Println("Login successfully saved to the environment")
	cfg.Passwod = enterValid("Password")
	err = os.Setenv("PASSWORD", cfg.Passwod)
	if err != nil {
		log.Println("Couldn't save password to the environment")
	}
	log.Println("Password successfully saved to the environment")

	// Email optional
	var email string
	fmt.Println("Please enter Email, will be helpfull if you lose your password")
	_, err = fmt.Scanln(&email)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	//
	// serverResp := RegistrationRespond{}
	serverResp, err := cfg.assistantClient.Register(cfg.Login, cfg.Passwod, email)
	if err != nil {
		return err
	}
	err = os.Setenv("USER_ID", serverResp.UserID)
	if err != nil {
		log.Println("Couldn't save UserID to the environment")
	}
	log.Println("UserID successfully saved to the environment")

	err = os.Setenv("MY_API_KEY", serverResp.APIKey)
	if err != nil {
		log.Println("Couldn't save Api key to the environment")
	}
	log.Println("Api key successfully saved to the environment")

	fmt.Println("Successfully registerd")
	fmt.Printf("UserID: %s\n", serverResp.UserID)
	fmt.Printf("Api Key: %s\n", serverResp.APIKey)
	return nil
}

func enterValid(s string) string {
	var temp string
	reader := bufio.NewReader(os.Stdin)

	for temp == "" {
		fmt.Printf("Enter your %s: \n", s)
		temp, _ = reader.ReadString('\n')
		temp = temp[:len(temp)-1]

		if temp == "" {
			fmt.Println("Input cannot be empty. Please try again.")
		}
	}

	return temp
}
