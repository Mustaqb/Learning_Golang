package main

import (
	"fmt"
	"strings"
)

func main() {
	i := 10
	ptr1 := &i
	pptr1 := &ptr1

	fmt.Printf("The value of i is %v and address is %p\n", i, &i)
	fmt.Println(strings.Repeat("#",20))
	fmt.Printf("The value of ptr1 is %v and address is %p\n", ptr1, &ptr1)
	fmt.Printf("The value where ptr1 is pointing is %v\n", *ptr1)
	fmt.Println(strings.Repeat("#",20))
	fmt.Printf("The value of pptr1 is %v and address is %p\n", pptr1, &pptr1)
	fmt.Printf("The value where pptr1 is pointing is %v\n", *pptr1)
	fmt.Printf("The value at the address where pptr1 is pointing is %v\n", **pptr1)
}
