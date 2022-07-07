package main

import "fmt"

func main() {
	var a int = 5

	//array values are intialized at the run time
	var arr =  [4]int{1,2,3,a};
	fmt.Printf("%v is of type %T of size %d\n",arr,arr,len(arr))

	var unint = [4]uint{1,2,1}
	fmt.Printf("%v is of type %T of size %d\n",unint,unint,len(unint))

	//Declaring array with unknown size using elippsis operator
	var array_ellipsis = [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("%v is of type %T of size %d\n",array_ellipsis,array_ellipsis,len(array_ellipsis))

	//This is slice operator
	var sl = []string{"a","b"}
	fmt.Printf("%v is of type %T of size %d\n",sl,sl,len(sl))
}