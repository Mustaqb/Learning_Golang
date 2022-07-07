package main

import (
	"fmt"
)

func main(){

	var (
		a int
		b int
	)

	a,b = 5,8
	fmt.Println(a,b,"Before swap")

	b,a=a,b
	fmt.Println(a,b,"After swap")
}