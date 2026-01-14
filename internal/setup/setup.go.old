package setup

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"dinhphu28.com/dictionary/internal/config"
	"dinhphu28.com/dictionary/internal/startup"
)

type RuntimeConfig struct {
	Mode          string `json:"mode"`
	ResourcesPath string `json:"resources_path"`
}

func Run() error {
	runtimeCfgPath, err := runtimeConfigPath()
	if err != nil {
		return err
	}

	cfgPath, err := configPath()
	if err != nil {
		return err
	}

	resPath, err := resourcesPath()
	if err != nil {
		return err
	}

	// Ensure dirs exist
	if err := os.MkdirAll(filepath.Dir(runtimeCfgPath), 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(cfgPath), 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(resPath, 0o755); err != nil {
		return err
	}

	runtimeCfg := RuntimeConfig{
		Mode:          "native",
		ResourcesPath: resPath,
	}

	predefinedCfgPath := startup.ResolvePath("config.json")
	if err := config.LoadConfig(predefinedCfgPath); err != nil {
		log.Fatal("failed to load config:", err)
	}
	cfg := config.GetGlobalConfig()

	if err := writeRuntimeConfig(runtimeCfgPath, runtimeCfg); err != nil {
		return err
	}

	if err := writeConfig(cfgPath, cfg); err != nil {
		return err
	}

	if err := installNativeManifest(); err != nil {
		return err
	}

	fmt.Println("✔ Dictionary setup complete")
	fmt.Println("Config:", runtimeCfgPath)
	fmt.Println("Resources:", resPath)
	fmt.Println("Run: dictionary")

	return nil
}

func writeRuntimeConfig(path string, cfg RuntimeConfig) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(cfg)
}

func writeConfig(path string, cfg config.GlobalConfig) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(cfg)
}

func runtimeConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "dictionary", "runtime.json"), nil
}

func configPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "dictionary", "config.json"), nil
}

func resourcesPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", "dictionary", "resources"), nil
}

func binPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, ".local", "bin", "dictionary"), nil
}
