package main

import "fmt"

func pass_by_value(a int) {
	a++
}

func pass_by_reference(a *int) {
	*a++
}

func main() {
	i := 10
	fmt.Printf("The value of i before pass by value %d\n", i)
	pass_by_value(i)
	fmt.Printf("The value of i after pass by value %d\n", i)
	pass_by_reference(&i)
	fmt.Printf("The value of i after pass by refence %d\n", i)

}
