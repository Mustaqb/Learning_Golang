package main

import (
	"fmt"
	"math"
)

type TwoDim interface {
	area()
}
type Square struct {
	side int
}

func (s Square) area() {
	fmt.Printf("The area of square of side %#v is %#v\n", s.side, s.side*s.side)
}

type ThreeDim interface {
	volume()
}
/*
	embedding TwoDim and ThreeDim in Object
*/
type Object interface {
	TwoDim
	ThreeDim
	getColor()
}
type Sphere struct{
	radius float64
	color string
}
func (s Sphere) volume() {
	fmt.Printf("The volume of sphere of radius %#v is %#v\n", s.radius, 4*math.Pi*s.radius*s.radius*s.radius/3)
}
func (s Sphere) area() {
	fmt.Printf("The area of sphere of radius %#v is %#v\n", s.radius, 4*math.Pi*s.radius*s.radius)
}
func (s Sphere) getColor(){
	fmt.Printf("The color of sphere of radius %#v is %#v\n", s.radius, s.color)
}

func main() {
	var square TwoDim= Square{
		side:10,
	}
	square.area()

	sphere:=Sphere{
		radius: 21,
		color: "Red",
	}
	sphere.area()
	sphere.volume()
	sphere.getColor()
}
