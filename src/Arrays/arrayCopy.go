package main

import "fmt"

func main(){
	arr:=[]int{1,2,3,4,5,6}

	fmt.Println(arr)

	arr_cpy:=make([]int,len(arr)-1)
	n := copy(arr_cpy,arr)

	
	fmt.Println(arr_cpy)
	fmt.Println(n)

	arr[2]=4
	fmt.Println(arr)
	fmt.Println(arr_cpy)


}