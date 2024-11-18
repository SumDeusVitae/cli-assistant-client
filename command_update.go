package main

import (
	"fmt"
	"os/exec"
	"regexp"

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
		fmt.Println(string(output))
		if err != nil {
			return err
		}
		command2 := exec.Command("sh", "-c", "mv ", string(output), "/bin/cli-assistant-client ", string(output), "/bin/qs")
		output2, err := command2.Output()
		fmt.Println("Output 2")
		fmt.Println(string(output2))
		if err != nil {
			return fmt.Errorf("command failed: %v\nOutput: %s", err, output)
		}
		// fmt.Println("You might need to run manually:")
		// fmt.Println("		mv $(go env GOPATH)/bin/cli-assistant-client $(go env GOPATH)/bin/qs")

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
