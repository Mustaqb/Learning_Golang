package main

import ("fmt")

func main()  {

	//creating new type
	type age int
	type new_age age
	var my_age age = 24
	var my_new_age new_age = 24
	fmt.Printf("%T and %v\n",my_age,my_age)
	fmt.Printf("%T and %v\n",my_new_age,my_new_age)


	//not creating new type, aliasing uint as uint2
	type uint2 = uint
	var i uint = 8
	var j uint2 = 16
	fmt.Println(j/i)
}