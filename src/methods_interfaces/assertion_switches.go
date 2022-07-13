package main

import "fmt"

type Numbers interface {
	Sum()
}

type Integers struct {
	value []int
}

//Integers implementing Numbers
func (i Integers) Sum() {
	sum := 0
	for v := range i.value {
		sum += i.value[v]
	}
	fmt.Printf("Sum of all integers is: %+v \n", sum)
}


type Decimals struct {
	value []float64
}

//Decimal implementing Numbers
func (d Decimals) Sum() {
	sum := 0.
	for v := range d.value {
		sum += d.value[v]
	}
	fmt.Printf("Sum of all decimals is: %+v \n", sum)
}

func (i Integers) Multiply() {
	prod := 1
	for v := range i.value {
		prod *= i.value[v]
	}
	fmt.Printf("Product of all integers is: %+v \n", prod)
}

func (d Decimals) count() {
	count := 0
	for v := range d.value {
		if d.value[v]-float64(int(d.value[v]))==0{
			count ++
		}
	}
	fmt.Printf("Count of integers in Decimal is: %+v \n", count)
}

func main() {
	/*
		Num is of type Numbers but holding value of Integers 
		since Integers implements Numbers
	*/
	var num Numbers=Integers{
		value:[]int{1,2,3,4,5,6,7},
	}
	fmt.Printf("The type of num is: %#v\n",num)
	/*
		Numbers variable(may they hold any other type data) 
		num can only call those functions that is defined in Numbers
	*/
	num.Sum()

	/*
		num cannot directly call any functions which are not defined in 
		Numbers but defined in the Integers
		to call Multiply using num we have to make use of assertion
	*/
	//we can call this way too but it is not safe
	fmt.Println("--------------Unsafe assertion num.(Integers).Multiply()--------------")
	num.(Integers).Multiply()

	//safe way to use assertion
	fmt.Println("--------------Safe assertion using safe variable--------------")
	num_int, safe:=num.(Integers)
	if safe{
		num_int.Multiply()
	}


	num = Decimals{
		value: []float64{1.1,2,3,4.4,5.5,6,7},
	}
	num.Sum()
	num_dec, safe:=num.(Decimals)
	if safe{
		num_dec.count()
	}

	/*
		Assertion using switch statement
	*/
	fmt.Println("--------------Asserting using Switch Statement--------------")
	switch value := num.(type) {
		case Integers:
			value.Multiply()
		case Decimals:
			value.count()
	}
}