package main

import (
	_ "flag"
	"fmt"
	"os"
)

func main() {
	test1()
	fmt.Println("------------------------- test2 -----")
	test2()
	fmt.Println("------------------------- test3 -----")
	test3()
	fmt.Println("------------------------- test4 -----")
	test4()
}

type Drawable interface {
	Draw(scale int) string
}

type Paintable interface {
	Paint(color int) string
}

type Diagram interface {
	Area() string
	Drawable
	Paintable
}

type Square struct {
	Size int
}

type Circle struct {
	Radius int
}

// *Square implement Draweable.Draw so it implements Drawable interface
func (s *Square) Draw(scale int) string {
	s.Size *= scale
	str := fmt.Sprintf("Draw Square: %v\tscale: %d\n", *s, scale)
	return str
}

// *Square implement Paintable.Paint so it implements Paintable interface
func (s Square) Paint(color int) string {
	str := fmt.Sprintf("Paint Square: %v\tcolor: %d\n", s, color)
	return str
}

// *square implement Diagram.Aread.
// and because implement Paintable.Paint and Drawable.Draw so Square is implemented Diagram
func (s *Square) Area() string {
	area := s.Size * s.Size
	str := fmt.Sprintf("Area Square: %v\tarea: %d\n", *s, area)
	return str
}

func (s *Square) String() string {
	return fmt.Sprintf("Square(%d)", s.Size)
}

// Circle implement Draweablr.Draw so it implements Drawable interface
func (c Circle) Draw(scale int) string {
	str := fmt.Sprintf("Draw Circle: %v\tscale: %d\n", c, scale)
	return str
}

// Circle implement Paintable.Paint so it implements Paintable interface
func (c Circle) Paint(color int) string {
	str := fmt.Sprintf("Paint Circle: %v\tcolor: %d\n", c, color)
	return str
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(%d)", c.Radius)
}

func test1() {
	var draw1 Drawable
	var Paint1 Paintable
	var diag1 Diagram

	s1 := Square{Size: 10}
	PrintValue(s1, "s1")
	c1 := Circle{Radius: 100}
	PrintValue(c1, "c1")

	// draw1 = s1	// ERROR: cannot use s1 (type Square) as type Drawable in assignment: Square does not implement Drawable (Draw method has pointer receiver)
	draw1 = &s1
	PrintValue(draw1, "draw1 = &s1")

	Paint1 = s1
	PrintValue(Paint1, "Paint1 = s1")
	Paint1 = &s1

	// diag1 = s1	// ERROR
	diag1 = &s1

	draw1 = c1
	draw1 = &c1
	Paint1 = c1
	Paint1 = &c1
	//diag1 = c1 	//ERROR
	//diag1 = &c1	//ERROR

	draw1 = diag1
	PrintValue(draw1, "draw1 = diag1")
	PrintValue(diag1, "diag1")
	Paint1 = diag1

	// draw1 = Paint1  //ERROR: cannot use Paint1 (type Paintable) as type Drawable in assignment: Paintable does not implement Drawable (missing Draw method)
	// diag1 = Paint1  //ERROR: cannot use Paint1 (type Paintable) as type Diagram in assignment: Paintable does not implement Diagram (missing Area method)
	draw1 = diag1
	Paint1 = diag1

	//-------
	draw1 = &s1
	draw1.Draw(2)
	//only methods of Drawable interface can be called on variable of Drawablr type.
	// draw1.Paint(3)  //ERROR: draw1.Paint undefined (type Drawable has no field or method Paint)

	//interface implementation assertion
	var _ Diagram = (*Square)(nil)
	// var _ Diagram = (*Circle)(nil) //ERROR: cannot use (*Circle)(nil) (type *Circle) as type Diagram in assignment: *Circle does not implement Diagram (missing Area method)
	// var _ Diagram = *new(Square) // ERROR : cannot use *new(Square) (type Square) as type Diagram in assignment: Square does not implement Diagram (Area method has pointer receiver)
	// var _ Diagram = *new(Circle) // ERROR: cannot use *new(Circle) (type Circle) as type Diagram in assignment: Circle does not implement Diagram (missing Area method)

	var _ Drawable = (*Square)(nil)
	var _ Drawable = (*Circle)(nil)
	//var _ Drawable = *new(Square) //ERROR: cannot use *new(Square) (type Square) as type Drawable in assignment: Square does not implement Drawable (Draw method has pointer receiver)
	var _ Drawable = *new(Circle)

	var _ Paintable = (*Square)(nil)
	var _ Paintable = (*Circle)(nil)
	var _ Paintable = *new(Square)
	var _ Paintable = *new(Circle)

	//---------------------------------
	var emptyInterface interface{} // empty interface match with every type
	emptyInterface = s1
	emptyInterface = &s1
	emptyInterface = c1
	emptyInterface = &c1
	emptyInterface = diag1
	emptyInterface = 5
	emptyInterface = int64(6)
	emptyInterface = os.Stdin
	_ = emptyInterface

	_ = draw1
	_ = Paint1
	_ = diag1
	_ = s1
	_ = c1
}

func test2() {

	// interface values ----------------------------------------------

	var di1 Diagram   //interface zero value is nil. both type and value now are nil
	di1 = &Square{10} // di1 type: *main.Square	value: Square{10}
	PrintValue(di1, "di1")
	di1 = nil            // type: nil	value: nil	di1 == nil
	di1 = (*Square)(nil) // type: */main.Square  value: nil		di1 != nil

	di1 = &Square{10} // type: *main.Square	value: Square{10}
	var pi1 Paintable // type: nil   value: nil
	pi1 = di1         //pi1 type: *main.Square	value: Square{10}
	//although now internal value of pi1 is *Square and *Square satisfy Diagram interface, but following statement is not valid; because pi1 itself
	// is of type Printable interface and Printable interface does'nt satisfy Diagram interface
	// di1 = pi1	//ERROR : cannot use pi1 (type Paintable) as type Diagram in assignment: Paintable does not implement Diagram (missing Area method)
	PrintValue(pi1, "pi1")

	_ = di1
	_ = pi1
}

func test3() {

	// type assertion for interfaces to concrete types

	s1 := Square{10}

	var dig1 Diagram
	dig1 = &s1

	//a1, ok := dig1.Square //ERROR: impossible type assertion: Square does not implement Diagram (Area method has pointer receiver)
	// a2, ok := dig1.(Circle) //ERROR: impossible type assertion: Circle does not implement Diagram (missing Area method)

	a1, ok := dig1.(*Square)
	fmt.Printf("%T\t%v\t%v\n", a1, a1, ok) // *main.Square    Square(10)      true

	var drw1 Drawable
	drw1 = &Square{10}
	//a2, ok := drw1.Square
	// drw1 now point to a *Square object, so its concrete type is equal to *Square and following statement returns true
	a2, ok := drw1.(*Square) // *main.Square    Square(10)      true
	fmt.Printf("%T\t%v\t%v\n", a2, a2, ok)
	// drw1 now point to a *Square object, so its concrete type is not equal to Circle or *Circle and following two statement returns flase
	a3, ok := drw1.(Circle)  // main.Circle     Circle(0)       false
	a4, ok := drw1.(*Circle) // *main.Circle    <nil>   		  false

	drw1 = Circle{20}
	a5, ok := drw1.(Circle)  //main.Circle     Circle(20)      true
	a6, ok := drw1.(*Square) // *main.Square    <nil>   false
	fmt.Printf("%T\t%v\t%v\n", a6, a6, ok)

	_, _, _, _ = a3, a4, a5, a6

}

func test4() {
	// type assertion for interfaces to interface types

	var drw1 Paintable = Square{10}
	a1, ok := drw1.(Diagram) // <nil>   <nil>   false
	fmt.Printf("%T\t%v\t%v\n", a1, a1, ok)
	// because dig1 point to an oject of type Square and Square does not satisfy Drawable interface, following statement result is false
	a2, ok := drw1.(Drawable) // <nil>   <nil>   false
	fmt.Printf("%T\t%v\t%v\n", a2, a2, ok)
	// because dig1 point to a oject of type Square and Square satisfy Paintable interface, following statement result is true
	a3, ok := drw1.(Paintable) // main.Square    Square(10)      true
	fmt.Printf("%T\t%v\t%v\n", a3, a3, ok)

	var p1 Paintable = Circle{20}
	// because Circle
	a4, ok := p1.(Diagram) // <nil>   <nil>   false
	fmt.Printf("%T\t%v\t%v\n", a4, a4, ok)
	a5, ok := p1.(Paintable) // main.Circle     Circle(20)      true
	fmt.Printf("%T\t%v\t%v\n", a5, a5, ok)

	_, _, _, _, _ = drw1, a1, a2, a3, a4
}

func PrintValue(a ...interface{}) {
	val := a[0]
	msg := ""
	for i := 1; i < len(a); i++ {
		msg += a[i].(string)
	}

	fmt.Printf("%s\ttype:%T\tvalue:%[2]v\n", msg, val)
}
