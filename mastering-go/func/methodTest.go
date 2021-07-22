package main

import (
	"fmt"
	"math"
)

func main() {
	test1()
	fmt.Println("---------------------------  test2 -------")
	test2()
	fmt.Println("---------------------------  test3 -------")
	test3()
}


// this method is belong to package main and dont belong to any type
func Distance(p, q Point) float64 {		//main.Distance
	fmt.Println("main.Distance")
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}



type Point struct {
	X, Y	float64
}

type Path []Point

type Path1 struct {
	points []Point
}

// p : method receiver. this method is defined on type Point
// no conflict between two following Distance method. below Distance is a method of type Point so is Point.Distance, 
// above Distance is a package level fumction: main.Distance (main is package name)
// Declaring a method with name X on Point => Error. because Point now has a field with name X
func (p Point) Distance(q Point) float64{	// Point.Distance
	fmt.Println("Point.Distance")
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}


// this method is not exported.
func (p Point) incrementBy(n float64) Point{
	p.X += n
	p.Y += n
	return p 
}


// when p is large and copy its value is not efficient or we need changes on p be visible for caller method, use pointer receiver
// golang convention: if any method of a type has Pointer receiver then all method of that type should have pointer receiver
// pointer receiver can be nil and this method is handled nil receiver 
// this method is exported
func (p *Point) ScaleBy(factor float64) {	// (*Point)).Distance
	if p == nil {
		return
	}
	p.X *= factor
	p.Y *= factor
}

// methods can be defined on any named type in the same package except types that have pointer or interface as their underlying type.
// each type has its own namespace so Point and Path can has its own Distance method. this two methods have different type
func (p Path) Distance(q Point) float64 {	// Path.distance
	fmt.Println("Path.Distance")
	return 0
}

func test1() {
	p1 := Point{1, 2}
	p2 := Point{4, 6}

	// p.Daistabce : selector
	fmt.Println(p1.Distance(p2))	// "Point.Distance"	Point.Distance isa called
	fmt.Println(Distance(p1, p2))	// "main.Distance" package level Distance is called

	perim := Path{{1,1}, {5,1}, {5,4}, {1,1}}
	fmt.Println(perim.Distance(p1))		// "Path.Distance"  Path.Distance is called

	// Path exposes its representation as a slice. so we can use both defined methods on it (e.g. Path.Distance) and also use slice operators and functions
	perim.Distance(p1)
	fmt.Println(perim[1])

	// Path1 dont expose its internal representation as slice. so only methods defined by this code can be called on them

}

func test2() {
	
	p1 := Point{1, 2}
	pp1 := &Point{3, 4}
	
	fmt.Println(p1) // {1,2}
	fmt.Println(p1.incrementBy(2))	// return value: {3,4} -- becuase incrementBy receiver is not pointer (Point), changes on receiver in incrementBy is not visible here on p1
	fmt.Println(p1) // {1,2}

	fmt.Println(pp1) // &{3,4}
	pp1.ScaleBy(2)	// becuase ScaleBy receiver is pointer to Point (*Point), changes on receiver in Scaleby is visible here on p3
	fmt.Println(pp1) // &{6,8}
	
	fmt.Println(p1) // {1,2}
	//Scaleby need pointer receiver and p1 is not pointer. Compiler automatically change statement to (&p1).ScaleBy
	p1.ScaleBy(2)
	fmt.Println(p1) //{2,4}

	fmt.Println(pp1) // &{6,8}
	// Scaleby need non pointer receiver and pp1 is pointer. Compiler automatically change statement to (*pp1).ScaleBy
	pp1.incrementBy(2)	
	fmt.Println(pp1) // &{6, 8}

	pp1 = nil
	// calling method on nil receiver is valid and receiver in ScaleBy have nil value. ScaleBy is handled nil value.
	pp1.ScaleBy(2)	
	fmt.Println(pp1) // nil
}


func test3() {
	p1 := Point{10,20}
	_ = p1
	
	// method values ---------------------------------
	f1 := p1.incrementBy		// p1.incrementBy returns a method value. f1 has method incrementBy on p 
	fmt.Printf("f1 type: %T\n", f1)	// func(float64) main.Point
	p2 := f1(3)	// f1(3) run p1.incrementBy(3)
	fmt.Println(p2)	// {13, 23}
	p2 = f1(5)
	fmt.Println(p2)	// {15, 25}


	f2 := p1.ScaleBy
	fmt.Printf("f2 type: %T\n", f2)	// func(float64)
	f2(4)	// run p1.ScaleBy(4)
	fmt.Println(p1)	// {40, 80}
	f2(2)	// run p1.ScaleBy(2)
	fmt.Println(p1)	// {80, 160}


	// method expressions ---------------------------
	
	p3 := Point{5, 6}
	f3 := Point.incrementBy	// method expression
	fmt.Printf("f3 type: %T\n", f3)	// func(main.Point, float64) main.Point
	p4 := f3(p3, 2)	// run p3.incrementBy(2)
	fmt.Println(p4)	// {7, 8}

	p5 := &Point{20, 30}
	f4 := (*Point).ScaleBy
	fmt.Printf("f4 type: %T\n", f4)	// func(*main.Point, float64)
	f4(p5, 2)	// run p5.ScaleBy(2)
	fmt.Println(p5)	// &{40, 60}


}