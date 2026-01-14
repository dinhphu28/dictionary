package setup

import "runtime"

type OSInfo struct {
	Name           string
	IsLinux        bool
	IsMac          bool
	IsWindows      bool
	Arch           string
	SupportsNative bool
}

func DetectOS() OSInfo {
	os := runtime.GOOS
	arch := runtime.GOARCH

	return OSInfo{
		Name:           os,
		IsLinux:        os == "linux",
		IsMac:          os == "darwin",
		IsWindows:      os == "windows",
		Arch:           arch,
		SupportsNative: os == "linux" || os == "darwin" || os == "windows",
	}
}
