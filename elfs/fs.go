package elfs

import (
	"github.com/el-ideal-ideas/ellib/elstr"
	"github.com/el-ideal-ideas/ellib/elsys"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	// sniff Length, use for detect file mime type
	MimeSniffLen = 512
)

var (
	// refer net/http package
	imageMimeTypes = map[string]string{
		"bmp":  "image/bmp",
		"gif":  "image/gif",
		"ief":  "image/ief",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"svg":  "image/svg+xml",
		"ico":  "image/x-icon",
		"webp": "image/webp",
	}
)

// PathExists reports whether the named file or directory exists.
func PathExists(path string) bool {
	if path == "" {
		return false
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// FileExists reports whether the named file or directory exists.
func FileExists(path string) bool {
	return IsFile(path)
}

// IsDir reports whether the named directory exists.
func IsDir(path string) bool {
	if path == "" {
		return false
	}
	if fi, err := os.Stat(path); err == nil {
		return fi.IsDir()
	}
	return false
}

// IsFile reports whether the named file or directory exists.
func IsFile(path string) bool {
	if path == "" {
		return false
	}
	if fi, err := os.Stat(path); err == nil {
		return !fi.IsDir()
	}
	return false
}

// MimeType get File Mime Type name. eg "image/png"
func MimeType(path string) (mime string) {
	if path == "" {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}

	return ReaderMimeType(file)
}

// ReaderMimeType get the io.Reader mimeType
// Usage:
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		return
// 	}
//	mime := ReaderMimeType(file)
func ReaderMimeType(r io.Reader) (mime string) {
	var buf [MimeSniffLen]byte
	n, _ := io.ReadFull(r, buf[:])
	if n == 0 {
		return
	}
	return http.DetectContentType(buf[:n])
}

// IsImageFile check file is image file.
func IsImageFile(path string) bool {
	mime := MimeType(path)
	if mime == "" {
		return false
	}
	for _, imgMime := range imageMimeTypes {
		if imgMime == mime {
			return true
		}
	}
	return false
}

// DeleteIfFileExist operate
func DeleteIfFileExist(filename string) error {
	if !IsFile(filename) {
		return nil
	}
	return os.Remove(filename)
}

// SelfDir returns The directory that contains the executable file.
// If current executable is an symbolic link,
// Return the real path to the directory that contains the executable.
func SelfDir() (path string, err error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	if exe, err = filepath.EvalSymlinks(exe); err != nil {
		return "", err
	}
	if exe, err = filepath.Abs(exe); err != nil {
		return "", err
	}
	return filepath.Dir(exe), nil
}

// GetTempDir returns the default directory to use for temporary files.
func GetTempDir() string {
	return os.TempDir()
}

// Get the path of parent directory.
func ParentDir(path string) (string, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	if elsys.IsWin() {
		i := strings.LastIndex(p, "\\")
		if i == -1 {
			return p, nil
		} else {
			return p[:i], nil
		}
	} else {
		i := strings.LastIndex(p, "/")
		if i == -1 {
			return p, nil
		} else if i == 0 {
			return "/", nil
		} else {
			return p[:i], nil
		}
	}
}

// Get filename extension
func GetFileExt(filename string) string {
	s := elstr.Split(filename, ".")
	if len(s) >= 2 {
		return s[len(s)-1]
	} else {
		return ""
	}
}

// Copy file
func CopyFile(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}

// Get the size of file as byte.
func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// Join File path list
func JoinMulti(path string, items ...string) string {
	if len(items) == 0 {
		return path
	} else {
		res := path
		for _, item := range items {
			res = Join(res, item)
		}
		return res
	}
}
