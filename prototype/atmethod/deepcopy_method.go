package main

import (
	"fmt"

	"github.com/pratapsgit/design-patterns-2024/prototype/student"
)

func main() {
	prat := student.Student{Name: "Prat",
		Location: &student.Address{Street: "7th croak street", City: "Bangalore", Pincode: "560087"},
		Friends:  []string{"Hello", "Challo"}}

	sam := prat

	//deep copy

	//sam.Location = prat.Location.DeepCopy()
	sam = *prat.DeepCopy()
	sam.Name = "Sam"
	sam.Location.Street = "6th croak stree"
	sam.Friends = append(sam.Friends, "Yellow")

	fmt.Println(prat, prat.Location)
	fmt.Println(sam, sam.Location)
}
