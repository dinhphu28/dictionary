package setup

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
)

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func detectBrowsers() (chrome, firefox bool) {
	if commandExists("google-chrome") ||
		commandExists("chromium") ||
		commandExists("chromium-browser") {
		chrome = true
	}

	if commandExists("firefox") {
		firefox = true
	}

	return
}

func chromeManifest(binaryPath, extID string) []byte {
	m := map[string]any{
		"name":        "com.dinhphu28.dictionary",
		"description": "Dictionary native host",
		"path":        binaryPath,
		"type":        "stdio",
		"allowed_origins": []string{
			"chrome-extension://" + extID + "/",
		},
	}
	b, _ := json.MarshalIndent(m, "", "  ")
	return b
}

func firefoxManifest(binaryPath string, extID string) []byte {
	m := map[string]any{
		"name":        "com.dinhphu28.dictionary",
		"description": "Dictionary native host",
		"path":        binaryPath,
		"type":        "stdio",
		"allowed_extensions": []string{
			extID,
		},
	}
	b, _ := json.MarshalIndent(m, "", "  ")
	return b
}

func installChromeManifest(data []byte) error {
	home, _ := os.UserHomeDir()
	dirs := []string{
		filepath.Join(home, ".config", "google-chrome", "NativeMessagingHosts"),
		filepath.Join(home, ".config", "chromium", "NativeMessagingHosts"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err == nil {
			path := filepath.Join(dir, "com.dinhphu28.dictionary.json")
			os.WriteFile(path, data, 0o644)
		}
	}
	return nil
}

func installFirefoxManifest(data []byte) error {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".mozilla", "native-messaging-hosts")
	os.MkdirAll(dir, 0o755)

	path := filepath.Join(dir, "com.dinhphu28.dictionary.json")
	return os.WriteFile(path, data, 0o644)
}
