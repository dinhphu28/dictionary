package setup

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type NativeManifest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Path        string   `json:"path"`
	Type        string   `json:"type"`
	Allowed     []string `json:"allowed_origins,omitempty"`
	Extensions  []string `json:"allowed_extensions,omitempty"`
}

func installNativeManifest() error {
	bin, err := os.Executable()
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	manifestDir := filepath.Join(
		home,
		".config",
		"google-chrome",
		"NativeMessagingHosts",
	)

	if err := os.MkdirAll(manifestDir, 0755); err != nil {
		return err
	}

	manifest := NativeManifest{
		Name:        "dictionary",
		Description: "Dictionary native host",
		Path:        bin,
		Type:        "stdio",
		Allowed: []string{
			"chrome-extension://kpgiaenkniiaacjbiipbmcdjfbjmgmll/",
		},
	}

	path := filepath.Join(manifestDir, "dictionary.json")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(manifest)
}
