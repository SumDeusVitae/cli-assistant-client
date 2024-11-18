package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/SumDeusVitae/cli-assistant-client/internal/version"
)

func callbackUpdate(cfg *config, args ...string) error {
	current := cfg.Variables.Version
	notLatest, err := version.CheckUpdate(current)
	if err != nil {
		return err
	}
	if notLatest {
		// install the latest version
		command := exec.Command("go", "install", "github.com/SumDeusVitae/cli-assistant-client@latest")
		_, err := command.Output()
		if err != nil {
			return err
		}
		command1 := exec.Command("go", "env", "GOPATH")
		output, err := command1.Output()
		if err != nil {
			return err
		}
		//
		string1 := strings.TrimSpace(string(output)) + "/bin/cli-assistant-client"
		string2 := strings.TrimSpace(string(output)) + "/bin/qs"
		command2 := exec.Command("mv", string1, string2)
		_, err = command2.Output()
		if err != nil {
			return fmt.Errorf("command failed: %v\nOutput: %s", err, output)
		}

		// Get the new version info
		command = exec.Command("qs", "version")
		b, err := command.Output()
		if err != nil {
			return err
		}
		re := regexp.MustCompile(`v\d+\.\d+\.\d+`)
		version := re.FindString(string(b))
		fmt.Printf("Successfully updated to %s\n", version)
	} else {
		fmt.Println("Latest version already installed!")
	}

	return nil

}
