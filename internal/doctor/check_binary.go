package doctor

import (
	"fmt"
	"os"
	"os/exec"
)

func checkBinary() {
	path, err := exec.LookPath("dictionary")
	if err != nil {
		fmt.Println("✖ dictionary binary not found in PATH")
		fmt.Println("  → Ensure ~/.local/bin is in PATH")
		return
	}

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("✖ dictionary binary not accessible:", err)
		return
	}

	fmt.Printf("✔ Binary: %s (%d KB)\n", path, info.Size()/1024)
}
