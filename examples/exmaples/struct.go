package examples

import (
	"fmt"

	"aheydari.ir/examples/utility"
)

type rectangle struct {
	length float64
	width  float64
	color  string
}

// A test
func A() {
	fmt.Println("Structs >>>>")

	var s1 rectangle
	utility.Describe(s1)

	s1.color = "red"
	utility.Describe(s1)

	var s2 = rectangle{1, 2, "green"}
	utility.Describe(s2)

	var s3 rectangle = rectangle{3, 4, "brown"}
	utility.Describe(s3)

	s3 = rectangle{length: 6}
	utility.Describe(s3)

	s3_1 := rectangle{4, 5, "pink"}
	utility.Describe(s3_1)

	var s4 = &rectangle{7, 8, "blue"}
	utility.Describe(s4)

	var s5 *rectangle = &rectangle{10, 11, "yellow"}
	utility.Describe(s5)

	var s6 *rectangle = new(rectangle)
	utility.Describe(s6)

	var s7 = new(rectangle)
	utility.Describe(s7)

	s8 := new(rectangle)
	utility.Describe(s8)

}
