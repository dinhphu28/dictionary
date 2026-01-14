package setup

import (
	"os"
	"path/filepath"
	"runtime"
)

type Paths struct {
	BinPath   string
	ConfigDir string
	DataDir   string
}

func DefaultPaths() Paths {
	return Paths{
		BinPath:   binaryInstallationPath(),
		ConfigDir: configDir(),
		DataDir:   dataDir(),
	}
}

func binaryInstallationPath() string {
	home, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "linux":
		return filepath.Join(home, ".local", "bin", "dictionary")
	case "darwin":
		return "/usr/local/bin/dictionary"
	case "windows":
		return filepath.Join(os.Getenv("LOCALAPPDATA"), "Dictionary", "dictionary.exe")
	default:
		panic("unsupported OS")
	}
}

func configDir() string {
	home, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "linux":
		return filepath.Join(home, ".config", "dictionary")
	case "darwin":
		return filepath.Join(home, "Library", "Application Support", "dictionary")
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "Dictionary")
	default:
		panic("unsupported OS")
	}
}

func dataDir() string {
	home, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "linux":
		return filepath.Join(home, ".local", "share", "dictionary")
	case "darwin":
		return filepath.Join(home, "Library", "Application Support", "dictionary")
	case "windows":
		return filepath.Join(os.Getenv("LOCALAPPDATA"), "Dictionary")
	default:
		panic("unsupported OS")
	}
}
