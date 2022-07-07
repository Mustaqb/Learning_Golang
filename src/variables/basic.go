package main

import (
	"fmt"
)

func main() {

	/*
		syntax is
		var <variable_name> <type>
		but type can be ommitted since the variable assignment happens from right to left so it knows
		before assignment that the var is string
	*/

	var name = "Mustaq"
	fmt.Println(name)
}
