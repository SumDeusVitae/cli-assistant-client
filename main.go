package main

import (
	"os"

	"github.com/SumDeusVitae/cli-assistant-client/internal/assistant"
)

type config struct {
	assistantClient assistant.Client
	Login           string
	Passwod         string
	Api             string
	UserID          string
}

func main() {
	cliClient := assistant.NewClient()
	cfg := &config{
		assistantClient: cliClient,
		Login:           os.Getenv("LOGIN"),
		Passwod:         os.Getenv("PASSWORD"),
		Api:             os.Getenv("MY_API_KEY"),
		UserID:          os.Getenv("USER_ID"),
	}
	runRep(cfg)
}
