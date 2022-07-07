package main

import (
	"fmt"
	"strings"
)

func main() {
	var sl = []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i:=0;i<len(sl);i++{
		fmt.Printf("at index %d of sl is %d\n", i, sl[i])
	}

	fmt.Println(strings.Repeat("#",20))

	for i, v := range sl {
		fmt.Printf("at index %d of sl is %d\n", i, v)
	}
}
