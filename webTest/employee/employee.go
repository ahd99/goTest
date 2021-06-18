package employee

import (
	"fmt"

	//"github.com/a-heydari/goutils/system"
	//"github.com/a-heydari/goutils/util"
	//"golang.org/x/tools/godoc/util"
)

type Salaried interface {
	getSalary() int
}

type Salary struct {
	basic 		int
	insurance	int
	allowance	int
}

type financial struct {
	amount 		int
	Salary
}

func (s Salary) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

type Employee1 struct {
	firstName, lastName		string
	isNew					bool
	salary					Salary
	uint16
	isalary					Salaried
}

func TestEmployee1() {
	var em1 Employee1	//initialize with zero values
	PrintValue(em1)

	var em2 Employee1 = Employee1{firstName: "ali", lastName: "heydari", isNew: true, salary: Salary{1,2,3}, uint16: 1345, isalary: Salary{}}  //inityialize when declare
	PrintValue(em2)
	PrintValue(em2.isalary)

	em3 := Employee1{"reza", "hasani", true, Salary{4,5,6}, 56, Salary{17, 18 ,19}}
	PrintValue(em3)
	
	em4 := struct {a string; b int} {"aaa", 12,}
	PrintValue(em4)

	var em5 struct {a int; b string} = struct{a int; b string}{13, "as13"}
	PrintValue(em5)
	
}

func TestEmployee2() {
	var emp1 *Employee1
	PrintValue(emp1)

	emp1 = &Employee1{}
	PrintValue(emp1)

	testStructPointer(emp1)


}

func TestEmployee3() {
	fin1 := financial{amount: 10000, Salary: Salary{2,3,4}}
	PrintValue(fin1)

	PrintValue(fin1.allowance)
}

func testStructPointer(e1 *Employee1) {
	PrintValue(e1)
}

func DoTest() {
	TestEmployee1()
	TestEmployee2()
	TestEmployee3()
}

func PrintValue(v interface{}) {
	fmt.Printf("%T >>  %v \n", v, v)
}
