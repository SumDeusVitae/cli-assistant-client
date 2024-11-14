package main

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/SumDeusVitae/cli-assistant-client/internal/version"
)

func callbackUpgrade(cfg *config, args ...string) error {
	current := cfg.Variables.Version
	notLatest, err := version.CheckUpdate(current)
	if err != nil {
		return err
	}
	if notLatest {
		// install the latest version
		command := exec.Command("go", "install", "github.com/sumdeusvitae/cli-assistant-client@latest")
		_, err := command.Output()
		if err != nil {
			return err
		}
	}
	// Get the new version info
	command := exec.Command("qs", "version")
	b, err := command.Output()
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`v\d+\.\d+\.\d+`)
	version := re.FindString(string(b))
	fmt.Printf("Successfully upgraded to %s!\n", version)

	return nil

}
