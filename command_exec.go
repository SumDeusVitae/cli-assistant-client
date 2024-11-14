package main

import (
	"fmt"

	"github.com/SumDeusVitae/cli-assistant-client/internal/version"
)

func callbackVer(cfg *config, args ...string) error {
	err := execute(cfg.Variables.Version)
	if err != nil {
		return err
	}
	return nil
}

func execute(ver string) error {
	info := version.FetchUpdateInfo(ver)
	defer info.PromptUpdateIfAvailable()
	fmt.Println(info)

	return nil
}
