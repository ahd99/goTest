package main

import (
	"a.heydari/test/carwler/pkg/requester"
	"github.com/a-heydari/goutils/util"
)
 
func main() {
	k := 5
	
	util.PrintValue(k)

	requester.Start()

}

