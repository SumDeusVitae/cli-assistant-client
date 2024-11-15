package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

const repoOwner = "sumdeusvitae"
const repoName = "cli-assistant-client"

type VersionInfo struct {
	CurrentVersion   string
	LatestVersion    string
	IsOutdated       bool
	IsUpdateRequired bool
	FailedToFetch    error
}

func FetchUpdateInfo(currentVersion string) VersionInfo {
	latest, err := getLatestVersion()
	if err != nil {
		return VersionInfo{
			FailedToFetch: err,
		}
	}

	isUpdateRequired := isUpdateRequired(currentVersion, latest)
	isOutdated := isOutdated(currentVersion, latest)
	return VersionInfo{
		IsUpdateRequired: isUpdateRequired,
		IsOutdated:       isOutdated,
		CurrentVersion:   currentVersion,
		LatestVersion:    latest,
	}
}

func CheckUpdate(current string) (bool, error) {
	latest, err := getLatestVersion()
	if err != nil {
		return false, err
	}
	isUpdateRequired := isOutdated(current, latest)
	return isUpdateRequired, nil

}

func CheckMajor(current string) (bool, error) {
	latest, err := getLatestVersion()
	if err != nil {
		return false, err
	}
	isUpdateRequired := isUpdateRequired(current, latest)
	return isUpdateRequired, nil
}

func (v *VersionInfo) PromptUpdateIfAvailable() {
	if v.IsOutdated {
		fmt.Fprintln(os.Stderr, "A new version of the CLI Assistant is available!")
		fmt.Fprintln(os.Stderr, "Please run the following command to update:")
		fmt.Fprintf(os.Stderr, "  qs update\n\n")
	}
}

// Returns true if the current version is older than the latest.
func isOutdated(current string, latest string) bool {
	return semver.Compare(current, latest) < 0
}

// Returns true if the latest version has a higher major or minor
// number than the current version. If you don't want to force
// an update, you can increment the patch number instead.
func isUpdateRequired(current string, latest string) bool {
	latestMajorMinor := semver.MajorMinor(latest)
	currentMajorMinor := semver.MajorMinor(current)
	return semver.Compare(currentMajorMinor, latestMajorMinor) < 0
}

func getLatestVersion() (string, error) {
	goproxyDefault := "https://proxy.golang.org"
	goproxy := goproxyDefault
	cmd := exec.Command("go", "env", "GOPROXY")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get GOPROXY: %v", err)
	}
	goproxy = strings.TrimSpace(string(output))

	proxies := strings.Split(goproxy, ",")
	if !contains(proxies, goproxyDefault) {
		proxies = append(proxies, goproxyDefault)
	}

	for _, proxy := range proxies {
		proxy = strings.TrimSpace(proxy)
		proxy = strings.TrimRight(proxy, "/")
		if proxy == "direct" || proxy == "off" {
			continue
		}
		modulePath := fmt.Sprintf("%s/%s", repoOwner, repoName)
		encodedModulePath := url.PathEscape(modulePath)
		url := fmt.Sprintf("%s/github.com/%s/@latest", proxy, encodedModulePath)
		// fmt.Printf("Trying proxy: %s\n", proxy)
		var resp *http.Response
		for retries := 0; retries < 3; retries++ {
			resp, err = http.Get(url)
			if err == nil {
				break
			}
			fmt.Printf("Error fetching %s (attempt %d): %v\n", url, retries+1, err)
			time.Sleep(2 * time.Second) // Wait before retrying
		}

		if err != nil {
			// All retries failed, move to the next proxy
			fmt.Printf("Failed to fetch from proxy %s after retries: %v\n", proxy, err)
			continue
		}

		/*
			resp, err := http.Get(url)
			if err != nil {
				continue
			}*/
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body from %s: %v\n", url, err)
			continue
		}
		// fmt.Printf("Response body from %s:\n%s\n", url, string(body))
		var ver struct{ Version string }
		if err = json.Unmarshal(body, &ver); err != nil {
			fmt.Printf("Error unmarshalling response from %s: %v\n", url, err)
			continue
		}
		// fmt.Printf("VERSION!!!!: %s\n", ver.Version)

		return ver.Version, nil
	}

	return "", fmt.Errorf("failed to fetch latest version")
}
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
