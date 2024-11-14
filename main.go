package main

import (
	_ "embed"
	"os"
	"strings"

	"github.com/SumDeusVitae/cli-assistant-client/internal/assistant"
	"github.com/SumDeusVitae/cli-assistant-client/internal/variables"
)

type config struct {
	assistantClient assistant.Client
	Variables       struct {
		Login    string
		Password string
		Api      string
		Version  string
	}
}

//go:embed version.txt
var versionString string

func main() {
	cliClient := assistant.NewClient()
	cfg := &config{
		assistantClient: cliClient,
	}
	cfg.Variables.Login = variables.LoadoadVariable("login")
	cfg.Variables.Password = variables.LoadoadVariable("password")
	cfg.Variables.Api = variables.LoadoadVariable("apiKey")
	// cfg.Variables.UserID = variables.LoadoadVariable("userId")
	err := callbackVer(cfg, strings.Trim(versionString, "\n"))
	if err != nil {
		os.Exit(1)
	}

	runRep(cfg)
}
