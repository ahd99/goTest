package main

import "log"

func main() {
	testByValue()
	testByPointer()
}

func testByValue() {
	log.Println("------- test by value")
	p1 := Person{"ali", 20}

	defer p1.StringV()     
	defer log.Println(p1)
	defer func() {
		p1.StringV()
	}()

	p1.name = "javad"

	defer p1.StringV()
	defer log.Println(p1)
	defer func() {
		p1.StringV()
	}()

	p1 = Person{"reza", 30}
}

func testByPointer() {
	log.Println("------- test by poiter")
	p1 := &Person{"ali", 20}

	defer p1.StringV()
	defer p1.StringP()
	defer log.Println(p1)
	defer func() {
		p1.StringV()
	}()

	p1.name = "javad"

	defer p1.StringP()
	defer log.Println(p1)
	defer func() {
		p1.StringV()
	}()
	
	p1 = &Person{"reza", 30}
}

type Person struct {
	name string
	age  int
}

func (p Person) StringV() {
	log.Println(p.name, p.age)
}

func (p *Person) StringP() {
	log.Println(p.name, p.age)
}