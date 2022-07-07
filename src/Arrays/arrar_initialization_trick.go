package main

import "fmt"

func main() {

	//we can also initialize only those index that are required
	//if used elipsis then the max index used will the size of the array
	//if no index provided then that element will take next element after the highest index used till then
	var arr = [...]int{
		99,
		6:101,
		3:1,
		4:93,
		100,
		10:90,
	}
	fmt.Printf("The array %#v is of size %d and type %T\n",arr,len(arr),arr)

	//99 has no index, previous element is placed at index 4 so it will take 5th index
	//100 has no index, previous index is at 4th so it will take 5th


	var arr23 = [...][6]int{
		3:{1,2,3,4,5},
		{5:'a',4:'d',2:'b','e'},
	}

	//b is at index 2 so e will take next to its previous element i.e 3 (next to 2)
	
	fmt.Printf("%v is of type %T of size %d inner of size %d\n",arr23,arr23,len(arr23),len(arr23[0]))

}