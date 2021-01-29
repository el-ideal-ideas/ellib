package elsys

import (
	"runtime"
	"syscall"
)

// GetMemoryUsage returns the size of memory used by this program as bytes.
func GetMemoryUsage() uint64 {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	return stat.Alloc
}

// GetDiskUsage returns the disk usage of given path as bytes.
func GetDiskUsage(path string) (used, free, total uint64) {
	fs := &syscall.Statfs_t{}
	err := syscall.Statfs(path, fs)
	if err == nil {
		total = fs.Blocks * uint64(fs.Bsize)
		free = fs.Bfree * uint64(fs.Bsize)
		used = total - free
	}
	return
}
