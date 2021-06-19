package util

import "fmt"

func PrintValue(v interface{}) {
	fmt.Printf("%T >>  %v \n", v, v)
}
