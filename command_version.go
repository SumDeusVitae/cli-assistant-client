package main

import (
	"fmt"

	"github.com/SumDeusVitae/cli-assistant-client/internal/version"
)

func callbackVer(cfg *config, args ...string) error {
	// fmt.Printf("CURRENT!!! %s\n", cfg.Variables.Version)
	info := check(cfg.Variables.Version)
	fmt.Println(info)
	if cfg.Variables.Outdated {
		outdated()
	}

	return nil
}

func check(ver string) string {
	info := version.FetchUpdateInfo(ver)
	defer info.PromptUpdateIfAvailable()
	return info.CurrentVersion
}

func outdated() error {
	fmt.Println("Update required")
	fmt.Println("QS CLI is outdated!")
	fmt.Println("please run: qs update")
	return nil
}
