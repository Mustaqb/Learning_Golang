package main

import (
	"fmt"
)

func main(){

	a:=16
	fmt.Printf("Printing Decimal %v to binary %b\n",a,a)
	fmt.Printf("Printing Decimal %v to binary %08b in 8 bit format\n",a,a)
	
	f1:=3.14
	f2:=2.56

	fmt.Printf("Multiplication of %v and %v is %.2f\n",f1,f2,f1*f2)
	fmt.Printf("Division of %v and %v is %.10f",f1,f2,f1/f2)

}