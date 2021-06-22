package hello

import (
	"fmt"
	"github.com/ahd99/goTest/goutils/util"
	"rsc.io/quote"
)

func Hello(s string) string {
	return fmt.Sprintf("Hello %v !", s)
}

func Hello1(s string) string {
	util.PrintValue(s)
	return s
}

func Hello2(s string) string {
	return quote.Hello()
}