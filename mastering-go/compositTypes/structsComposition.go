package main

import "fmt"

func main() {
	test1()
	test2()
	test3()
}

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point //struct composition
	Radius int
}

type Wheel struct {
	Circle // anonymous field: a filed of Another type without name. Circle is embedded within Wheel
	Spokes int
}

type Wheel1 struct {
	*Circle // Pointer anonymous field: a filed point to Another type without name. *Circle is embedded within Wheel
	Spokes  int
}

func test1() {
	var c Circle

	c.Center.X = 10
	c.Center.Y = 20
	c.Radius = 5
	c.Center = Point{8, 9}

	//literal form1: all fields values without name
	c1 := Circle{Point{5, 6}, 20}
	//literal form2 name:value format
	c1 = Circle{Radius: 10, Center: Point{5, 6}}
	c1 = Circle{Center: Point{5, 6}}

	fmt.Printf("%v\t%v\n", c, c1)
}

func test2() {
	var w Wheel

	w.Spokes = 5

	w.Circle.Radius = 6 //Circle is anonymous field in Wheel so we use its type as name for accessing its fields
	w.Radius = 10       // using Circle name is optional for accessing its fields because it is anonnumous filed n Wheel struct
	w.Circle = Circle{Center: Point{1, 2}, Radius: 20}

	w.Circle.Center.X = 7
	w.Center.X = 11

	w.Circle.Center.Y = 8
	w.Center.Y = 12

	// w.X = 13 	// compile error. because Point is not anonymous field in Circle struct and has a name Center

	// in literal all struct names must be use. no shorthand is possible
	//literal format 1
	w1 := Wheel{Circle{Point{5, 6}, 20}, 30}
	//literal format 2
	w1 = Wheel{Spokes: 30, Circle: Circle{Center: Point{5, 6}, Radius: 20}}

	//it is nor possible to have two anonyn=mous field of the same type because of name conflict

	// suppose circle type is defined like circle{X,Y int} (circle itself with lowercase first letter but its fields are exported)
	//and Wheel define it as anonymous field. in another packages (w is of type Wheel):
	// w.circle.X	raise compile error because circle name is not exported
	// but w.X is ok and X is accessible from other packages.

	fmt.Printf("%#v\n%#v\n", w, w1)
}

func test3() {
	var w Wheel1

	w.Spokes = 5

	w.Circle = new(Circle) //because *Circle is anonymous field.
	w.Circle.Radius = 6    // *Circle is anonymous field in Wheel1 so we use its type as name for accessing its fields
	w.Radius = 10          // using Circle name is optional for accessing its fields because it is anonnumous filed n Wheel1 struct
	w.Circle = &Circle{Center: Point{1, 2}, Radius: 20}

	w.Circle.Center.X = 7
	w.Center.X = 11

	w.Circle.Center.Y = 8
	w.Center.Y = 12

	// w.X = 13 	// compile error. because Point is not anonymous field in Circle struct and has a name Center

	// in literal all struct names must be use. no shorthand is possible
	//literal format 1
	w1 := Wheel1{&Circle{Point{5, 6}, 20}, 30}
	//literal format 2
	w1 = Wheel1{Spokes: 30, Circle: &Circle{Center: Point{5, 6}, Radius: 20}}

	//it is nor possible to have two anonyn=mous field of the same type because of name conflict

	// suppose circle type is defined like circle{X,Y int} (circle itself with lowercase first letter but its fields are exported)
	//and Wheel define it as anonymous field. in another packages (w is of type Wheel):
	// w.circle.X	raise compile error because circle name is not exported
	// but w.X is ok and X is accessible from other packages.

	fmt.Printf("%#v\n%#v\n", w, w1)
}
