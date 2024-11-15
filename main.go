package main

import (
	_ "embed"
	"strings"

	"os"

	"github.com/SumDeusVitae/cli-assistant-client/internal/assistant"
	"github.com/SumDeusVitae/cli-assistant-client/internal/variables"
	"github.com/SumDeusVitae/cli-assistant-client/internal/version"
)

type config struct {
	assistantClient assistant.Client
	Variables       struct {
		Login    string
		Password string
		Api      string
		Version  string
		Outdated bool
	}
}

//go:embed version.txt
var ver string

func main() {
	cliClient := assistant.NewClient()
	cfg := &config{
		assistantClient: cliClient,
	}
	cfg.Variables.Login = variables.LoadoadVariable("login")
	cfg.Variables.Password = variables.LoadoadVariable("password")
	cfg.Variables.Api = variables.LoadoadVariable("apiKey")
	// cfg.Variables.UserID = variables.LoadoadVariable("userId")
	cfg.Variables.Version = strings.Trim(ver, "\n")
	// fmt.Printf("version.txt context: %s\n", strings.Trim(ver, "\n"))
	check, err := version.CheckMajor(cfg.Variables.Version)
	if err != nil {
		os.Exit(1)
	}
	if check {
		cfg.Variables.Outdated = true
	}

	runRep(cfg)
}
