package main

import (
	"fmt"

	mathi "aheydari.ir/test/packages/math"
	stri "aheydari.ir/test/packages/string"
	//stri "aheydari.ir/test/packages/string"
)

func main() {
	fmt.Println("Hello packages test")
	fmt.Println(mathi.Add(1, 4))
	fmt.Println(stri.Greetings("ali"))

}
