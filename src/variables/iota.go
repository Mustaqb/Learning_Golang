package main

import (
	"fmt"
)

func main(){
	const (
		a = iota
		b
		c
	)
	fmt.Println(a,b,c)

	//print 2k+1
	const (
		x = iota*2 + 1
		y
		z
	)

	fmt.Println(x,y,z)
	
	//print skipping 3
	const (
		u=iota
		v
		w
		_
		p
	)
	fmt.Println(u,v,w,p)
}