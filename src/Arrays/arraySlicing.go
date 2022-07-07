package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	arr1 := arr[1:4]

	fmt.Println(arr,arr1)

	arr1=append(arr1,6)

	fmt.Println(arr,arr1)
}
