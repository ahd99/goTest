package main

import (
	"hello/square"
	"log"
	"math"
)

//import "rsc.io/quote"

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//test1()
	//test2()

	res := square.SquareRoot(129)
	log.Print("sqr root res:", res, "    ", math.Pow(res, 2))

	log.Print("Finish")

}
