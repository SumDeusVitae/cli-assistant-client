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
		fmt.Println("Output 1")
		if err != nil {
			return err
		}
		string1 := strings.TrimSpace(string(output)) + "/bin/cli-assistant-client"
		fmt.Println("String1: ", string1)
		string2 := strings.TrimSpace(string(output)) + "/bin/qs"
		fmt.Println("String2: ", string2)
		// uploading v0.3.0 for testing purposes

		command2 := exec.Command("mv", string1, string2)
		output2, err := command2.Output()
		fmt.Println("Output 2")
		fmt.Println(string(output2))
		if err != nil {
			fmt.Printf("command failed: %v\nOutput: %s\n", err, output)
			fmt.Println("You might need to manually run:")
			fmt.Println("		mv $(go env GOPATH)/bin/cli-assistant-client $(go env GOPATH)/bin/qs")
			return err
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
