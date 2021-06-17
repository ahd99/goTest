package utility

import (
	"fmt"
	"reflect"
)

// Describe def
func Describe(i interface{}) {
	fmt.Printf("%v, %T %v \n", i, i, reflect.TypeOf(i))
}
