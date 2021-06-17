package main

import (
	"fmt"
	"log"
	"aheydari.ir/greetings"
)

func callGreetings() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	//fmt.Println("Hello, World!")
	//fmt.Println(quote.Go())
	message, err := greetings.Hello("ali")
	if( err != nil) {
		log.Fatal(err)
	}
	fmt.Println(message)
}