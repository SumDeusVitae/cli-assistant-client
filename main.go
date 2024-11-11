package main

import (
	"github.com/SumDeusVitae/cli-assistant-client/internal/assistant"
	"github.com/SumDeusVitae/cli-assistant-client/internal/variables"
)

type config struct {
	assistantClient assistant.Client
	Variables       struct {
		Login    string
		Password string
		Api      string
		UserID   string
	}
}

func main() {
	cliClient := assistant.NewClient()
	cfg := &config{
		assistantClient: cliClient,
	}
	cfg.Variables.Login = variables.LoadoadVariable("login")
	cfg.Variables.Password = variables.LoadoadVariable("password")
	cfg.Variables.Api = variables.LoadoadVariable("apiKey")
	cfg.Variables.UserID = variables.LoadoadVariable("userId")

	runRep(cfg)
}
