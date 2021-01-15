package env

import (
	"os"
	"runtime"
)


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

// GetEnv get environment value by key name with default value.
func GetEnv(name string, def ...string) string {
	val := os.Getenv(name)
	if val == "" && len(def) > 0 {
		val = def[0]
	}
	return val
}