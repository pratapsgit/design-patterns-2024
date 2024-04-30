package main

import (
	"fmt"

	"github.com/pratapsgit/design-patterns-2024/prototype/student"
)

func main() {
	prat := student.Student{Name: "Prat", Location: &student.Address{Street: "7th croak street", City: "Bangalore", Pincode: "560087"}}

	sam := prat

	//deep copy

	sam.Location = &student.Address{
		Street:  prat.Location.Street,
		City:    prat.Location.City,
		Pincode: prat.Location.Pincode}

	sam.Name = "Sam"
	sam.Location.Street = "6th croak stree"

	fmt.Println(prat, prat.Location)
	fmt.Println(sam, sam.Location)
}
