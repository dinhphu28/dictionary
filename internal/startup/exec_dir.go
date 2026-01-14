package startup

import (
	"os"
	"path/filepath"
)

func ResolvePath(name string) string {
	if p := os.Getenv("DICT_BASE"); p != "" {
		return filepath.Join(p, name)
	}
	return filepath.Join(exeDir(), name)
}

func exeDir() string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exe)
}
