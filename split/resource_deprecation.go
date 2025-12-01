package split

import (
	"os"
)

// isHarnessTokenSet checks if the harness_token or harness_platform_api_key is set in the provider configuration
// or in the environment variable
func isHarnessTokenSet(meta interface{}) bool {
	// First check if meta is available and contains harness_token or harness_platform_api_key
	if meta != nil {
		if config, ok := meta.(*Config); ok && config.harnessToken != "" {
			return true
		}
	}

	// Fallback to environment variable if meta isn't available during plan
	harnessPlatformAPIKey := os.Getenv("HARNESS_PLATFORM_API_KEY")
	harnessToken := os.Getenv("HARNESS_TOKEN")
	return harnessPlatformAPIKey != "" || harnessToken != ""
}
