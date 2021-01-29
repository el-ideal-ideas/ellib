package elsys

import "runtime"

// If running on windows, return true.
func IsWin() bool {
	return runtime.GOOS == "windows"
}

// If running on mac, return true.
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// If running on linux, return true.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}
