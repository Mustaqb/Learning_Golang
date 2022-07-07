package main

import (
	"fmt"
)

func main(){

	var a="Mustaq"
	/*
	the type of the variable is assigned and known at the compile time. so, go completes the code by
	adding the type at compile time
	at compile time a is known to store string so it is assigned type string
	*/
	// a=10
	/*
	Assignment happens at the run time so a is given 10 which throws error.
	Go is a Statically and Strong Typed Programming Language
	*/
	fmt.Println(a)


	var (
		price float64
		name string
		age int
		alive bool
	)

	print(price,name,age,alive)
}