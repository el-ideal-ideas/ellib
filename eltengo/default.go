package eltengo

import (
	"fmt"
	"runtime"
)

var defaultScript = fmt.Sprintf(`
/* The Tengo Language */

goVersion := "%s"

fmt := import("fmt")

each := func(seq, fn) {
    for x in seq { fn(x) }
}

`, runtime.Version())
