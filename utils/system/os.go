package system

import "runtime"

// IsMac 是否为Mac
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsWindows 是否为Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux 是否为Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}
