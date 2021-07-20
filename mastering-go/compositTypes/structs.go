package main

import "fmt"

func main() {
	test1()
	test2()
}

//struct Employee can not contain field of the same type Employee (recursion). but it can have a field with pointer to itself: *Employee

type Employee struct {
	ID          int
	Name        string
	Age, Salary int
	managerID   int //this field not exported because first letter is lowercase, so is not accessible outside of current package
	parent      *Employee
}

func test1() {
	var eme1 Employee // structs are value type. so eme1 is not nil and is a Employee that all fields have zero-value
	fmt.Println(eme1) //zero value of struct is composed of zero value of all its fields: {0 "" "" 0 0 0 <nil>}

	type emptyStruc struct{} //emptystruct. zero size, no data

	//struct literal 1 : wihtout fields name and only values. need all fields in struct definition order
	eme2 := Employee{2001, "ali", 30, 30000, 1001, nil}
	fmt.Println(eme2)

	//struct literal 2: with fields name and corresponding values. if a field name ommited in literal, it get zero-value. order is not important
	eme5 := Employee{Age: 30, ID: 2001, Name: "ali"} // other fields get zero-value (Acording their type)
	fmt.Println(eme5)

	var em1 Employee
	em1.ID = 1001
	em1.Name = "ali"
	sname := em1.Name
	fmt.Println(em1.ID, sname)

	emp1 := &em1
	(*emp1).Age = 30
	emp1.Age = 35

	spname := &em1.Salary
	*spname = 20000

	EmployeeByID(1002).Name = "noName"

	fmt.Println(em1)

	// if all fileds of a struct are comaprable, the struct is comparable so two expression of that type can be compared by == or !=
	// two struct are equal if all corresponding fields are equal
	if em1 == eme1 {
		//equal
	}

}

func test2() {
	em1 := Employee{ID: 1001}

	// fields names, types and order is part of type. em1 is type of Employee
	// following struct type and Employee is one the same type because name, type and order of fields or same each other.
	// we can write whole type each time it is needed or use its name: Employee
	em1 = struct {
		ID        int
		Name      string
		Age       int
		Salary    int
		managerID int
		parent    *Employee
	}{ID: 1002}

	fmt.Println(em1)

	em1 = Employee{ID: 1003}

}

func EmployeeByID(id int) *Employee {
	em := Employee{}
	em.ID = id
	return &em
}
