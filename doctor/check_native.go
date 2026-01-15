package doctor

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func checkNativeMessaging() {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		fmt.Println("ℹ Native messaging not supported on this OS")
		return
	}

	chrome := filepath.Join(
		os.Getenv("HOME"),
		".config/google-chrome/NativeMessagingHosts",
		"com.dinhphu28.dictionary.json",
	)

	firefox := filepath.Join(
		os.Getenv("HOME"),
		".mozilla/native-messaging-hosts",
		"com.dinhphu28.dictionary.json",
	)

	found := false

	if _, err := os.Stat(chrome); err == nil {
		fmt.Println("✔ Chrome native messaging installed")
		found = true
	}

	if _, err := os.Stat(firefox); err == nil {
		fmt.Println("✔ Firefox native messaging installed")
		found = true
	}

	if !found {
		fmt.Println("⚠ Native messaging not installed")
		fmt.Println("  → Run dictionary setup")
	}
}
