package config

import (
	"fmt"
	"net/http"
	"strings"
)

// Checks to make sure all required environment variables are set
func validateRequired(config *Config) error {
	if config.Hostname == "" {
		return fmt.Errorf("%s is required", HostnameEnv)
	}

	if config.TSControlURL == "" {
		return fmt.Errorf("%s is required", TSControlURLEnv)
	}

	if config.HPControlURL == "" {
		return fmt.Errorf("%s is required", HPControlURLEnv)
	}

	if config.TSAuthKey == "" {
		return fmt.Errorf("%s is required", TSAuthKeyEnv)
	}

	if config.HPAuthKey == "" {
		return fmt.Errorf("%s is required", HPAuthKeyEnv)
	}

	return nil
}

// Pings the Tailscale control server to make sure it's up and running
func validateTSReady(config *Config) error {
	testURL := config.TSControlURL
	if strings.HasSuffix(testURL, "/") {
		testURL = testURL[:len(testURL)-1]
	}

	// TODO: Consequences of switching to /health (headscale only)
	testURL = fmt.Sprintf("%s/key?v=109", testURL)
	resp, err := http.Get(testURL)
	if err != nil {
		return fmt.Errorf("Failed to connect to TS control server: %s", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to connect to TS control server: %s", resp.Status)
	}

	return nil
}

// Pings the Headplane server to make sure it's up and running
func validateHPReady(config *Config) error {
	testURL := config.HPControlURL
	if strings.HasSuffix(testURL, "/") {
		testURL = testURL[:len(testURL)-1]
	}

	testURL = fmt.Sprintf("%s/healthz", testURL)
	resp, err := http.Get(testURL)
	if err != nil {
		return fmt.Errorf("Failed to connect to HP control server: %s", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to connect to HP control server: %s", resp.Status)
	}

	return nil
}
