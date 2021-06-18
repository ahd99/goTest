package main

import (
	"log"
	"os"
	"strconv"
)

func main()  {
	
	args := os.Args
	if len(args) < 2 {
		log.Println("No Arguments Error.")
		return
	}

	var sum float64 = 0
	for i:= 1; i < len(args); i++ {
		n, err := convert(args[i])
		if err == nil {
			sum += n
		}
	}

	log.Println("sum:", sum)
	
}

func convert(s string) (float64, error) {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
			log.Println(s, "is not a valid float and ignored.")
	}
	return n, err
}