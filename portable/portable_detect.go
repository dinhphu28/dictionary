package portable

import (
	"log"
	"os"

	"github.com/dinhphu28/dictionary/internal/startup"
)

func IsPortable() bool {
	return hasConfig()
}

func hasConfig() bool {
	configPath := startup.ResolvePath("config.json")
	configPathToml := startup.ResolvePath("config.toml")
	log.Printf("CONFIG PATH: %v | %v", configPath, configPathToml)
	return fileExists(configPath) || fileExists(configPathToml)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
