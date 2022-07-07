package main

import (
	"fmt"
)

func main(){

	/*
	constants are created and initialized at compile time
	*/
	const a=10
	fmt.Println(a)
	fmt.Printf("Type of a is %T\n",a)

	var v1 = 10
	fmt.Printf("Type of v1 is %T\n",v1)
	
	/*
	constants should always be initialized with constant values, i.e constants cannot be 
		initialized using variable
		because variables get their values at runtime and constants need their value at compile-time
	*/
	var i=10
	const j=i

	const (
		const1=100
		const2
		const3
	)

	fmt.Println(const1,const2,const3)


	//Variable operations are done at the runtime so divide by zero is not known till runtime
	v1,v2:=5,0
	fmt.Print(v1/v2)

	//constants operations are done at the compile-time so error is known divide by zero before running code
	const c1,c2=4,0
	fmt.Print(c1/c2)

	fmt.Print(c1/v2)
	fmt.Print(v1/c2)


}