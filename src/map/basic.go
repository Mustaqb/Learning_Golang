package main

import "fmt"

func main() {
	var m map[int]string
	
	fmt.Printf("%#v\n",m)


	m2:=map[int]string{}
	m2[0]="123"
	fmt.Printf("%#v\n",m2)

	m1:=make(map[int]string)
	m1[1]="NEW"
	fmt.Printf("%#v\n",m1)

}
