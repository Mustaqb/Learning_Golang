package main

import (
	"fmt"
)

func main()  {
	i:=0
outer:
	if i<5{

		i++
		fmt.Println(i)
		goto outer
	}
}