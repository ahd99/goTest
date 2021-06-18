package main

import (
	"log"
	"os"
)

func main() {
	//use default logger that write to stderr.default logger is defined as a global variable (std) in log package
	// var std = New(os.Stderr, "", LstdFlags)
	log.Println("hello")
	
	// create a cusgtom logger that write to stdout 
	// with differnt flgs
	var logger1 *log.Logger
	logger1 =  log.New(os.Stdout, "mylog ", log.Ldate | log.Ltime | log.Lshortfile)
	logger1.Println("My log in stdout")

	// ope (or create) a file for writing log to it.
	var logFile string = "/Users/a.heydari/workspace/gotest/log/test1.log"
	var f * os.File
	f, err := os.OpenFile(logFile, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error openning file ", logFile, err)
	}
	defer f.Close()

	// create a custm logger that write to file
	logger2 := log.New(f, "mfl ", log.Ldate | log.Ltime )
	logger2.Println("do you see in file?")

	//change flags in logger
	logger2.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	logger2.Println("file lof flags changes")
}