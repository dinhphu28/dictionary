package setup

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Install(paths Paths) error {
	// 1. Install binary
	binPath := paths.BinPath
	if err := installBinary(binPath); err != nil {
		return err
	}

	if err := installConfigs(paths); err != nil {
		return err
	}

	// 2. Install resources
	resourcesPath := filepath.Join(paths.DataDir, "resources")
	if err := installResources(resourcesPath); err != nil {
		return err
	}

	return nil
}

func installBinary(path string) error {
	if err := copyFile(
		"./dictionary",
		path,
		0o755,
	); err != nil {
		return fmt.Errorf("install binary: %w", err)
	}
	return nil
}

func installConfigs(paths Paths) error {
	// 1. Install config.toml (only if not exists)
	// cfgDst := filepath.Join(paths.ConfigDir, config.ConfigFile)
	// if _, err := os.Stat(cfgDst); os.IsNotExist(err) {
	// 	if err := copyFile(startup.ResolvePath(config.ConfigFile), cfgDst, 0o644); err != nil {
	// 		return err
	// 	}
	// }

	WriteConfigToml(paths.ConfigDir, DefaultConfig(), false)

	return nil
}

func InstallNativeMessagingManifests(bin string) {
	chrome, firefox := detectBrowsers()

	if chrome {
		manifest := chromeManifest(bin, "kpgiaenkniiaacjbiipbmcdjfbjmgmll")
		if err := installChromeManifest(manifest); err != nil {
			log.Fatalf("install chrome native messaging manifest failed: %v", err)
		}
	}

	if firefox {
		manifest := firefoxManifest(bin, "503e78dec27c89515afd99f62ecf12e3305a204d@temporary-addon")
		if err := installFirefoxManifest(manifest); err != nil {
			log.Fatalf("install firefox native messaging manifest failed: %v", err)
		}
	}
}

func installResources(resourcesPath string) error {
	if _, err := os.Stat("./resources"); err == nil {
		if err := copyDir("./resources", resourcesPath); err != nil {
			return err
		}
	}
	return nil
}
