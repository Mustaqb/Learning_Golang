package main

import "fmt"

type student struct{
	name string
	marks []int
	percentage float64
}

func (s *student) calc_percentage(){
	total_marks:=0
	for i := range s.marks{
		total_marks+=i
	}
	s.percentage=float64(total_marks/len(s.marks))
}

func main() {
	mustaq:=student{
		name:"Mustaq",
		marks:[]int{100,86,98,92,94},
	}

	mustaq.calc_percentage()
	fmt.Printf("%#v\n",mustaq)
}