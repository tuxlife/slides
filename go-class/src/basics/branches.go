package main

import (
	"fmt"
	"runtime"
)

func main() {
	const str string = "this is a 日本語 string\n"
	if len(str) > 0 {
		fmt.Printf("str: %q\n", str)
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%q.", os)
	}
}
