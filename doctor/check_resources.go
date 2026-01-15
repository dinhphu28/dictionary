package doctor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dinhphu28/dictionary/database"
)

func checkResources() {
	resDir := filepath.Join(
		os.Getenv("HOME"),
		".local/share/dictionary/resources",
	)

	if _, err := os.Stat(resDir); err != nil {
		fmt.Println("✖ Resources directory missing:", resDir)
		fmt.Println("  → Copy dictionaries to this directory")
		return
	}

	if err := database.LoadDictionaries(resDir); err != nil {
		fmt.Println("✖ Failed to load dictionaries:", err)
		return
	}

	dicts := database.GetDictionaries()
	fmt.Printf("✔ Dictionaries loaded: %d\n", len(dicts))

	for _, d := range dicts {
		fmt.Printf("  - %s (%s)\n", d.Manifest.FullName, d.Manifest.ID)
	}
}
