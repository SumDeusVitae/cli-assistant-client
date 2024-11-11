package variables

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Get the file path where the API key will be stored
func getVariableFilePath(variableType string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, ".cli-assistant", variableType)
}

// Save the API key to the file
func SaveVariable(variableType, v string) error {
	variableFile := getVariableFilePath(variableType)

	// Create the directory if it does not exist
	dir := filepath.Dir(variableFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Save the variable in the file using os.WriteFile
	err := os.WriteFile(variableFile, []byte(v), 0600) // only readable/writable by the user
	if err != nil {
		return err
	}
	return nil
}

// Load the saved API key
func LoadoadVariable(variableType string) string {
	variableFile := getVariableFilePath(variableType)
	data, err := os.ReadFile(variableFile)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}
