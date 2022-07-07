package main

import "fmt"

func main() {

	/*
	each element of string is byte type, byte type basically supports
	small range of characters

	Rune is another type of char element storing type which supports
	wide range of characters

	by default char declared are rune type
	*/
	var a byte = 'a'
	fmt.Printf("%T %v %c\n",a,a,a)
	var b rune ='b'
	fmt.Printf("%T %v %c\n",b,b,b)

	fmt.Println(rune(a)+b)
	fmt.Println((a)+byte(b))
}
