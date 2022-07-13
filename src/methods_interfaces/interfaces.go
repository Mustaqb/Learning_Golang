package main

import (
	"fmt"
	"math"
)

/* 
	If any Struct(Rectangle/Circle) wants to implement interface(Shape) 
	then, Struct has to implement all the functions that is declared in
	interface i.e if the Rectangle/Circle is implementing perimeter_area 
	then it is implementing Shape.

	Note: Struct can still have functions other than those mentioned
	in Interface, like Rectangle has is_square
*/
type Shape interface {
	perimeter_area()
}

type Rectangle struct {
	length    int
	breadth   int
	area      int
	perimeter int
}

type Circle struct {
	radius    int
	area      float64
	perimeter float64
}

func (circle *Circle) perimeter_area() {
	circle.perimeter = math.Pi * float64(2*circle.radius)
	circle.area = math.Pi * float64(math.Pow(float64(circle.radius), 2))
}

func (rectangle *Rectangle) perimeter_area() {
	rectangle.perimeter = 2 * (rectangle.length + rectangle.breadth)
	rectangle.area = rectangle.length * rectangle.breadth
}

func (rectangle Rectangle)is_square(){
	if rectangle.length==rectangle.breadth{
		fmt.Printf("The rectangle is Square\n")
	}else{
		fmt.Printf("The rectangle is not Square\n")
	}
}

func show_shape_info(shape Shape) {
	shape.perimeter_area()
	fmt.Printf("Circle Information: %+v\n", shape)
}

func main(){
	circle:=Circle{
		radius: 7,
	}
	show_shape_info(&circle)
	

	rectangle:=Rectangle{
		length: 10,
		breadth: 20,
	}
	show_shape_info(&rectangle)
	rectangle.is_square()
}