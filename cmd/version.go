package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

// Placeholder for builds
var Version = "1.0.0"

type Release struct {
	TagName string `json:"tag_name"`
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the CLI version",
	Long:    `This command prints the CLI version`,
	Example: "scrapPY version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scrapPY: " + Version)
		latestReleaseTag, err := latestRelease()
		if err == nil && latestReleaseTag != "" {
			latestRelease := strings.TrimPrefix(latestReleaseTag, "v")
			currentRelease := strings.TrimPrefix(Version, "v")
			if semver.Compare(latestRelease, currentRelease) > 0 {
				updateScrapPY(latestRelease)
			}
		}
	},
}

// Fetch latest release for comparison to current version
func latestRelease() (string, error) {
	resp, err := http.Get("https://api.github.com/repos/RoseSecurity/ScrapPY/releases/latest")
	if err != nil {
		return "", fmt.Errorf("failed to fetch version: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var release Release
	if err := json.Unmarshal(body, &release); err != nil {
		return "", fmt.Errorf("failed to parse version: %w", err)
	}

	return release.TagName, nil
}

// Display out of date warning
func updateScrapPY(latestVersion string) {
	c1 := color.New(color.FgCyan)

	c1.Println(fmt.Sprintf("\nYour version of ScrapPY is out of date. The latest version is %s\n\n", latestVersion))
}
