package system

import (
	"fmt"
	"runtime"
)

// SystemInfo return System info
func SystemInfo() string {
	return fmt.Sprintf("Goroutins v4: %v\n", runtime.NumGoroutine())
	//return "System Info !"
}
