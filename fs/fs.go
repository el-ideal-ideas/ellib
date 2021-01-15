package fs

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)


const (
	// sniff Length, use for detect file mime type
	MimeSniffLen = 512
)

var (
	// refer net/http package
	imageMimeTypes = map[string]string{
		"bmp": "image/bmp",
		"gif": "image/gif",
		"ief": "image/ief",
		"jpg": "image/jpeg",
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

// The directory of the executable file.
func SelfDir() (path string, err error) {
	exe, err := os.Executable()
	if err != nil {
		return
	}
	return filepath.Dir(exe), nil
}
