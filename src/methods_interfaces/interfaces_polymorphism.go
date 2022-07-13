package main

import "fmt"

type Numbers interface {
	Addition()
}

type Integers struct {
	value []int
}

type Float struct {
	value []float64
}

func (i Integers) Addition() {
	sum := 0
	for v := range i.value {
		sum += i.value[v]
	}
	fmt.Printf("Sum of all integers is: %+v \n", sum)
}

func (f Float) Addition() {
	sum := 0.
	for v := range f.value {
		sum += float64(f.value[v])
	}
	fmt.Printf("Sum of all integers is: %+v \n", sum)
}

func main() {
	var num Numbers
	fmt.Printf("The type of num is: %#v\n",num)


	integer:=Integers{
		value:[]int{1,2,3,4,5,6,7},
	}
	integer.Addition()
	//integer is implementing Number so we can assign integer to num
	num=integer
	fmt.Printf("The type of num is: %#v\n",num)

	floating:=Float{
		value:[]float64{1,2,3,4,5,6,7},
	}
	floating.Addition()
	//Since Float also implements Number, dynamically we can also change
	//the assignment from Integers to Float
	num=floating
	fmt.Printf("The type of num is: %#v\n",floating)

}
