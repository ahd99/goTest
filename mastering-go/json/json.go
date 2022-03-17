package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	test1()
}

func test1() {
	b, err := json.Marshal(1234)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("%s\n", b)

	var i interface{}
	err = json.Unmarshal(b, i)
	fmt.Printf("%d\n", i)

}
