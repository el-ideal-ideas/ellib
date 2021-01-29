package elcgo

/*
#include <stdlib.h>
extern void hello();
extern void message(char* message);
*/
import "C"
import "unsafe"

// Print Hello World use cgo.
func Hello() {
	C.hello()
}

// Print message use cgo.
func Message(msg string) {
	cStr := C.CString(msg)
	defer C.free(unsafe.Pointer(cStr))
	C.message(cStr)
}
