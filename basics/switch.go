package basics

import (
	"fmt"
	"runtime"
)

// Go only runs the selected case, not all the cases that follow.
// In effect, the break statement that is needed at the end of each case
// in those languages is provided automatically in Go.
// Another important difference is that Go's switch cases need not be
// constants, and the values involved need not be integers.
func Switch() {
	// switch os := runtime.GOOS; os { case statements...}
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// Switch without a condition is the same as switch true.
