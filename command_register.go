package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SumDeusVitae/cli-assistant-client/internal/variables"
)

func callbackRegister(cfg *config, args ...string) error {
	// Login and Password
	cfg.Variables.Login = enterValid("Login")
	// log.Println("Login successfully saved to the environment")
	cfg.Variables.Password = enterValid("Password")

	// log.Println("Password successfully saved to the environment")

	// Email optional
	email := ""
	fmt.Println("Please enter Email, will be helpfull if you lose your password")
	reader := bufio.NewReader(os.Stdin)
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	//

	serverResp, err := cfg.assistantClient.Register(cfg.Variables.Login, cfg.Variables.Password, email)
	if err != nil {
		return err
	}
	fmt.Printf("Successfully registerd: %s\n", cfg.Variables.Login)
	// SAVING VARIABLES TO LOCAL

	// fmt.Printf("Api Key: %s\n", serverResp.APIKey)
	// saved to cfg
	cfg.Variables.Api = serverResp.APIKey
	err = variables.SaveVariable("apiKey", serverResp.APIKey)
	if err != nil {
		log.Println("Couldn't save api to local variable")
	}
	//
	/*
		fmt.Printf("UserID: %s\n", serverResp.ID)
		cfg.Variables.UserID = serverResp.ID
		err = variables.SaveVariable("userId", serverResp.ID)
		if err != nil {
			log.Println("Couldn't save user id to local variable")
		}
	*/
	//
	err = variables.SaveVariable("login", cfg.Variables.Login)
	if err != nil {
		log.Println("Couldn't save login to local variable")
	}

	err = variables.SaveVariable("password", cfg.Variables.Password)
	if err != nil {
		log.Println("Couldn't save password to local variable")
	}

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
