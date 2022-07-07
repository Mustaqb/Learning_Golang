package main

import (
	"fmt"
)

func main()  {
	i := 10
	ptr := &i
	fmt.Printf("The pointer ptr is pointing at address %p which holds value %d\n",ptr,*ptr)

	//decalaring pointer this way will not point anywhere so can't use *ptr1
	var ptr1 *int
	fmt.Printf("The pointer ptr1 is pointing at address %p\n",ptr1)

	//declaring pointer using new will point to open space with value 0
	var ptr2 = new(int)
	fmt.Printf("The pointer ptr2 is pointing at address %p which holds value %d\n",ptr2,*ptr2)
	ptr3 := new(int)
	fmt.Printf("The pointer ptr3 is pointing at address %p which holds value %d\n",ptr3,*ptr3)

}