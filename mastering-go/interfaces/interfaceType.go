package main

import "fmt"

func main() {
	test1()
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
	var paint1 Paintable
	var diag1 Diagram

	s1 := Square{Size: 10}
	c1 := Circle{Radius: 100}

	// draw1 = s1	// ERROR: cannot use s1 (type Square) as type Drawable in assignment: Square does not implement Drawable (Draw method has pointer receiver)
	 // diag1 = s1	// ERROR : cannot use s1 (type Square) as type Diagram in assignment: Square does not implement Diagram (Area method has pointer receiver)
	
	paint1 = s1	
	draw1 = &s1
	paint1 = s1
	paint1 = &s1
	diag1 = &s1

	// paint1 = draw1	// cannot use draw1 (type Drawable) as type Paintable in assignment: Drawable does not implement Paintable (missing Paint method)
	// diag1 = paint1	//cannot use paint1 (type Paintable) as type Diagram in assignment: Paintable does not implement Diagram (missing Area method)
	paint1 = diag1
	draw1 = diag1

	draw1 = c1
	draw1 = &c1
	paint1 = c1
	paint1 = &c1

	_ = draw1
	_ = paint1
	_ = diag1
	_ = s1
	_ = c1
}
