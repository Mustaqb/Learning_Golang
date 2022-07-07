package main

import "fmt"

func update_age_by_reference(stu *student) {
	stu.age = 24
}

func update_age_by_value(stu student) {
	stu.age = 24
}

type student struct {
	name string
	age  int
}

func main() {
	mustaq := student{
		name: "Mustaq",
	}

	//values in struct are only changed using pointers
	fmt.Println("Before calling pass by value", mustaq)
	update_age_by_value(mustaq)
	fmt.Println("After calling pass by value", mustaq)
	update_age_by_reference(&mustaq)
	fmt.Println("After calling pass by reference", mustaq)
}
