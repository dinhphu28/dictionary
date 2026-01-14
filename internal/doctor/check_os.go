package doctor

import (
	"fmt"
	"runtime"
)

func checkOS() {
	fmt.Printf("✔ OS: %s (%s)\n", runtime.GOOS, runtime.GOARCH)
}
