package main

import (
	"log"
	"time"
)

var logger *log.Logger = log.Default()

func main()  {
	do()
}

func do() {
	s := "job1"
	defer trace(s)() //trace(s) will be executed in this line and "Trace start" will be printed. Anonymous method returned from trace(s) method will be executed at end of do() method
	time.Sleep(5 * time.Second)
}

func trace(s string) func() {
	start := time.Now()
	logger.Println(" Trace strt -", s)
	return func() {
		logger.Println(" Trace finish -", s, "  ", time.Since(start))
	}
} 

