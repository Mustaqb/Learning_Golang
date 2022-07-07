package main

import "fmt"

func update_array_by_reference(arr [5]int) {
	arr[2] = 30
}

func update_slice_by_reference(arr []int) {
	arr[2] = 30
}

func update_map_by_refernce(m map[int]int){
	m[5]=5
}

func main() {
	//declaring array
	arr := [5]int{1, 2, 3, 4, 5}
	update_array_by_reference(arr)
	fmt.Printf("Array after update treated as array %#v\n", arr)


	//To modify array pass the array as slice

	/*
		logic:-
		when slice is created from array, the slice just points to the 
		array (backing array), no extra space is created for slice, so 
		any changes to the slice effects the array when we pass the 
		array as slice then any changes to slice effects the array
	*/
	update_slice_by_reference(arr[:])
	fmt.Printf("Array after update treated as slice %#v\n", arr)
	sl:=[]int{1,2,3,4,5}
	update_slice_by_reference(sl)
	fmt.Printf("Slice after update %#v\n", sl)

	//map variable holds address of the location where data is kept
	m:=map[int]int{0:0,1:1,2:2,3:3,4:4}
	fmt.Printf("Map before update %#v\n", m)
	update_map_by_refernce(m)
	fmt.Printf("Map after update %#v\n", m)
}
