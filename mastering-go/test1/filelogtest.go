package main

import (
	"log"
	"os"
)

func main() {

	// use default logger in log package (std gobal variable)
	// var std = New(os.Stderr, "", LstdFlags)
	// write to stderr (stderr goes to terminal as default )
	// empty prefix
	// LstdFlags: default configuration (LstdFlags     = Ldate | Ltime)
	log.Println("hello ")

	//create another logger to stdout
	var logger1 *log.Logger = log.New(os.Stdout, "MyLog ", log.Ldate|log.Ltime|log.Lshortfile)
	logger1.Println("salam")

	var LOGFILE string = "~/goTest/logs/logTest.tmp"

	var f *os.File
	var err error
	f, err = os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Panic("can not open log file")
	}
	defer f.Close()

	//var logger2 *log.Logger = log.New()

}
