package eldev

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"
)

// GetFuncName get func name
func GetFuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// Get the line number that this function was called.
func GetCurrentLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

// Get the file name that this function was called.
func GetCurrentFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// Get the line number that the caller of this function was called.
func GetCallerLine() int {
	_, _, line, _ := runtime.Caller(2)
	return line
}

// Get the file name that the caller of this function was called.
func GetCallerFile() string {
	_, file, _, _ := runtime.Caller(2)
	return file
}

// Get the function name who call the caller of this function.
func GetCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	return f.Name()
}

// PkgName get current package name
func GetPkgName() string {
	_, filePath, _, _ := runtime.Caller(1)
	file, _ := os.Open(filePath)
	r := bufio.NewReader(file)
	line, _, _ := r.ReadLine()
	pkgName := bytes.TrimPrefix(line, []byte("package "))
	return string(pkgName)
}

// PanicIfErr if error is not empty
func PanicIfErr(err error) {
	if err != nil {
		pc, _, line, _ := runtime.Caller(2)
		name := runtime.FuncForPC(pc).Name()
		panic(fmt.Sprintf("[Panic: <%d - %s>] %s", line, name, err))
	}
}

// dump the goroutine stack information.
func DumpStacks(depth ...uint) {
	var buf []byte
	if len(depth) == 0 {
		buf = make([]byte, 16384)
	} else {
		buf = make([]byte, depth[0])
	}
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}

// GetCallStacks stacks is a wrapper for runtime.
// Stack that attempts to recover the data for all goroutines.
func GetCallStacks(all bool) []byte {
	// We don't know how big the traces are, so grow a few times if they don't fit.
	// Start large, though.
	n := 10000
	if all {
		n = 100000
	}
	var trace []byte
	for i := 0; i < 5; i++ {
		trace = make([]byte, n)
		bts := runtime.Stack(trace, all)
		if bts < len(trace) {
			return trace[:bts]
		}
		n *= 2
	}
	return trace
}

// Return the type of value as string.
func GetType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// Return the time to execute function `f`
// Similar to Python's timeit.timeit
func Timeit(f func(), number int) time.Duration {
	start := time.Now()
	for i := 0; i < number; i++ {
		f()
	}
	return time.Since(start)
}
