package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(
		runtime.Compiler,
		runtime.GOARCH,
		runtime.Version(),	// go version
		runtime.NumCPU(),		
		runtime.NumGoroutine(),
	)
}