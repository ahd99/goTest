package main

import (
	"errors"
	"fmt"
)

func main()  {
	fmt.Println("in main")
	err := func1()
	fmt.Println("main finish. err: ", err) //run
}

func func1() (e error) {
	defer func() {		
		if p := recover(); p != nil {
			fmt.Println("in recover f1 - ", p)
			e = errors.New(fmt.Sprintf("panic error. msg: %v", p))
		}
	} ()

	fmt.Println("func1 start")
	func2()
	fmt.Println("func1 finish") //dont run
	return nil
}

func func2() {
	defer func() { fmt.Printf("func2 defer") }()
	fmt.Println("func2 start")
	panic("<< panic f2 >>")
	fmt.Println("func2 finish") // dont run
}

func func3() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.New(fmt.Sprintf("panic error. msg: %v", p))
		}
	} ()

	// calling panic or do anything that results to a panic, stops execution of this function but function returns normally with an error and caller function dont sense the panic
	// if someSituations {
	// 	panic("panic error message")
	// }

	return nil
}