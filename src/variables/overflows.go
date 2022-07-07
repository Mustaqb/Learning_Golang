package main

import (
	"fmt"
	"math"
)

func main(){

	//127 is added to i in compile time, i can hold max value of 8 bit i.e 127
	var i int8 = 127
	fmt.Println(i,"is printed")

	//128 is added at compile time so go knows int8 type var store max value of 128
	// var j int8 = 128
	// fmt.Println(j,"is not printed")

	//when i is incremented it goes to -128 its lowest
	fmt.Println(i+1,"Reaches to lowest")

	//Flost stored max value
	var f float32 = (math.MaxFloat32)
	fmt.Println(f,"is printed")

	//when float is incremented
	fmt.Println(f*1.2,"INF reached")

}