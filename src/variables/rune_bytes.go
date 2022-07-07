package main

import (
	"fmt"
)

func main(){

	// single char uses single quotes and string uses double quotes
	// rune and byte are used for single char

	//Throws error as string cannot be used to assign rune
	// var ru rune="M"
	// fmt.Println(ru)

	//This is correct
	var b byte='~'
	var r rune='~'
	fmt.Printf("b = %c with type %T\n",b,b)
	fmt.Printf("r = %c with type %T\n",r,r)

	//Throws error as rune cannot be used to assign string
	// var str string='m'
	// fmt.Println(str)

	var str1 string ="M"
	fmt.Println(str1)

	var new_rune ='r'
	fmt.Printf("Type of %c is %T\n",new_rune,new_rune)

	var new_str ="R"
	fmt.Printf("Type of %c is %T\n",new_str,new_str)
}