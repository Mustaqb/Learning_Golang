package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	r := "世界k"

	fmt.Println("Rune count in string 世界k is",utf8.RuneCountInString(r))
	fmt.Println("Length of string 世界k is",len(r))

	fmt.Println("Printing Rune char by char")
	for i:=0;i<len(r);i++{
		fmt.Printf("%c",r[i])
	}
}
