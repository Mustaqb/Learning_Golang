package main

import (
	"fmt"
)

func main(){

	//Array type has strict size type, replaced with 0 if less alue provided
	var arr = [4]int{1,2,3,4}
	fmt.Printf("%v is of type %T\n",arr,arr)

	var arr1 = [5]int{1,2,3,4}
	fmt.Printf("%v is of type %T\n",arr1,arr1)

	//Slice has dynamic data type
	var sl = []int{1,2,3,4,5,6,7}
	fmt.Printf("%v is of type %T\n",sl,sl)

	//Map is key value pair syntax map[key_type]value_type
	var m = map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	fmt.Printf("%v is of type %T\n",m,m)
	fmt.Println("'a' in m holds",m["a"])

	//Struct is to store different data-type var together
	type Person struct {
		age int
		name string
		weight int
	}

	var mustaq Person
	fmt.Printf("%T,%v\n",mustaq,mustaq)

	//Pointer holds address of another variable
	var a="M"
	var ptr = &a
	fmt.Printf("ptr stores address %v of variable %T\n",ptr,ptr)

	//Function is also a type of variable
	fmt.Printf("%v is of type %T\n",fun,fun)
}

func fun(){

}