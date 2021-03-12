package elsys

import (
	"os"
	"runtime"
	"strings"
)

var isRHEL bool
var isRHEL8 bool
var isRHELLike bool
var isCentOS bool
var isCentOS8 bool

func init() {
	if redhat, err := os.ReadFile("/etc/redhat-release"); err == nil {
		contents := string(redhat)
		isRHEL = strings.HasPrefix(contents, "Red Hat Enterprise Linux release")
		isRHEL8 = strings.HasPrefix(contents, "Red Hat Enterprise Linux release 8")
		isRHELLike = true
		isCentOS = strings.HasPrefix(contents, "CentOS Linux release")
		isCentOS8 = strings.HasPrefix(contents, "CentOS Linux release 8")
	} else {
		isRHEL = false
		isRHEL8 = false
		isRHELLike = false
		isCentOS = false
		isCentOS8 = false
	}
}

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

// If running on freebsd, return true.
func IsFreeBSD() bool {
	return runtime.GOOS == "freebsd"
}

// If running on unix-like system, return true.
func IsUnixLike() bool {
	return runtime.GOOS == "darwin" || runtime.GOOS == "linux" || runtime.GOOS == "freebsd"
}

// If running on Red Hat Enterprise Linux, return true.
func IsRHEL() bool {
	return isRHEL
}

// If running on Red Hat Enterprise Linux 8, return true.
func IsRHEL8() bool {
	return isRHEL8
}

// If running on CentOS, return true.
func IsCentOS() bool {
	return isCentOS
}

// If running on CentOS 8, return true.
func IsCentOS8() bool {
	return isCentOS8
}

// If running on linux that based on RHEL, return true.
func IsRHELLike() bool {
	return isRHELLike
}
