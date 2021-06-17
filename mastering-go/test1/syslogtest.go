package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {

	log.Println("stdErr LLLogggg !!") //write log to stderr (default log destinaion)

	//craete sys log:
	sysLog, err := syslog.New(syslog.LOG_LOCAL7|syslog.LOG_INFO, "myApp")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	defer log.Println("do you see defer message?")

	log.Println("log_MAIL - LOG_INFO !!!!")
	//log.Fatal("ooooh fatal !!!")
	log.Panic("oooooh panic !!!")

	log.Println("Do you see after fatal log command?")
	fmt.Println("Do you see after fatal fmt print command?")
}
