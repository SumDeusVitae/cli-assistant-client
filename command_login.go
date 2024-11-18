package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SumDeusVitae/cli-assistant-client/internal/variables"
)

func callbackLogin(cfg *config, args ...string) error {
	if cfg.Variables.Outdated {
		outdated()
		return nil
	}
	login, password := "", ""
	// Checking if API already assigned
	if cfg.Variables.Api != "" {
		fmt.Println("Seems like you are already logged in are you sure that you want to login?")
		fmt.Print("Do you want to continue? [Y/n]: ")
		// Create a reader to capture input
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)
		switch input {
		case "", "y":
			fmt.Println("OK")
		case "n":
			return nil
		default:
			fmt.Println("Invalid input. Please enter 'Y' or 'n'.")
			return nil
		}

	}

	if cfg.Variables.Login != "" && cfg.Variables.Password != "" {
		fmt.Println("Seems like you have saved Login & Password")
		fmt.Println("Do you wish to use them? [Y/n]")
		// Create reader
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" || strings.ToLower(input) == "y" {
			login = cfg.Variables.Login
			password = cfg.Variables.Password
		}
	}
	for login == "" {
		fmt.Print("Please enter Login: ")
		login, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		login = login[:len(login)-1]

		if login == "" {
			fmt.Println("Input cannot be empty. Please try again.")
		}
	}
	for password == "" {
		fmt.Print("Please enter password: ")
		password, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		password = password[:len(password)-1]

		if password == "" {
			fmt.Println("Input cannot be empty. Please try again.")
		}
	}
	// saving locally for future
	cfg.Variables.Login = login
	err := variables.SaveVariable("login", login)
	if err != nil {
		log.Println("Couldn't save user login to local variable")
	}
	cfg.Variables.Password = password
	err = variables.SaveVariable("passowd", password)
	if err != nil {
		log.Println("Couldn't save user password to local variable")
	}
	serverResp, err := cfg.assistantClient.Login(login, password)
	if err != nil {
		return err
	}
	fmt.Println("Logged In Successfully!")
	// SAVE API KEY locally
	cfg.Variables.Api = serverResp.APIKey
	err = variables.SaveVariable("apiKey", serverResp.APIKey)
	if err != nil {
		log.Println("Couldn't save api to local variable")
	}

	return nil
}
