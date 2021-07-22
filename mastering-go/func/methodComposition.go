package main

import "fmt"

func main() {
	test1()

}

type Point struct {
	X, Y float64
}

// Point is embedded in ColoredPoint => Composition
type ColoredPoint struct{
	Point
	color	int
}

// *Point is embedded in ColoredPoint => Composition
type ColoredPoint_P struct{
	*Point
	color	int
}

func (p Point) incrementBy(n float64) Point{
	p.X += n
	p.Y += n
	return p
}


func (p *Point) ScaleBy(factor float64) {
	if p == nil {
		return
	}
	p.X *= factor
	p.Y *= factor
}


func test1() {

	c1 := ColoredPoint{Point{1,2}, 10}

	// Point is embedded in ColoredPoint so we can access to Point fields directly from ColoredPoint
	n1 := c1.Point.X
	n1 = c1.X
	c1.X = 3
	_ = n1
	
	fmt.Println(c1)	//{{3 2} 10}
	// because Point is embedded in ColoredPoint, there us access to Point methods from ColoredPoint 
	// the methods of Point have been promoted to ColoredPoint
	p1 := c1.incrementBy(2)
	fmt.Println(c1, p1)	//{{3 2} 10} {5 4}
	c1.ScaleBy(3)
	fmt.Println(c1)	//{{3 2} 10} {5 4}

	/*
	when compiler resolve resolve a selector such as p.ScaleBy to its method (means search to find ScaleBy), it first
	looks for a directly declared method name ScalBy, then for methods promoted once from directly embedded types, then for methods 
	promoted twice from embedded types and so on.
	if two method with same name are promoted from the same level => compile ERROR
	*/

	// also when pointer to a type is embedded, method and field promotion works.
	cp := ColoredPoint_P{&Point{1,2,}, 10}
	fmt.Println(cp, cp.Point)	//{0xc0000ac030 10}  &{1,2}
	n1 = cp.Point.X
	n1 = cp.X
	cp.incrementBy(2)
	fmt.Println(cp, cp.Point)	//{0xc0000ac030 10}  &{1,2}
	cp.ScaleBy(3)
	fmt.Println(cp, cp.Point)	//{0xc0000ac030 10}  &{3,6}

}