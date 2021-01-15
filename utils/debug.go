package utils

import (
	"bufio"
	"bytes"
	"os"
	"reflect"
	"runtime"
)


// FuncName get func name
func FuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// PkgName get current package name
func PkgName() string {
	_, filePath, _, _ := runtime.Caller(1)
	file, _ := os.Open(filePath)
	r := bufio.NewReader(file)
	line, _, _ := r.ReadLine()
	pkgName := bytes.TrimPrefix(line, []byte("package "))
	return string(pkgName)
}
