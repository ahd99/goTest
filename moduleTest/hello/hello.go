package hello

import (
	"fmt"
	"github.com/ahd99/goTest/goutils/util"
)

func Hello(s string) string {
	return fmt.Sprintf("Hello %v !", s)
}

func Hello1(s string) string {
	util.PrintValue(s)
	return s
}