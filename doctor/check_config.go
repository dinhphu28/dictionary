package doctor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dinhphu28/dictionary/internal/config"
)

func checkConfig() {
	cfgPath := filepath.Join(
		os.Getenv("HOME"),
		".config/dictionary/config.json",
	)

	if _, err := os.Stat(cfgPath); err != nil {
		fmt.Println("✖ Config missing:", cfgPath)
		fmt.Println("  → Run dictionary setup")
		return
	}

	if err := config.LoadConfig(cfgPath); err != nil {
		fmt.Println("✖ Config invalid:", err)
		return
	}

	cfg := config.GetGlobalConfig()
	fmt.Printf("✔ Config loaded (%d priorities)\n", len(cfg.Priority))
}
