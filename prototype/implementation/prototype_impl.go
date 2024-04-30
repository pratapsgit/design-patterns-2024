package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Addr struct {
	Unit   int32
	Block  string
	Street string
	City   string
}

type Employee struct {
	Name     string
	EType    string
	Location *Addr
}

func (e *Employee) DeepCopy() *Employee {
	//create a buffer
	b := bytes.Buffer{}
	//create a encoder pass the buffer
	en := gob.NewEncoder(&b)
	//encode the required object
	_ = en.Encode(e)

	//create a decoder
	dc := gob.NewDecoder(&b)
	//create an empty object
	res := Employee{}
	//decode the buffer to empty object
	_ = dc.Decode(&res)

	return &res
}

// create 2 prototypes for employees
var permEmployee = &Employee{"", "Permanent", &Addr{0, "E block", "Laurel street", "Delhi"}}

var contrEmployee = &Employee{"", "Contract", &Addr{0, "F block", "Laurel street", "Delhi"}}

func newEmployee(etype *Employee, name string, unit int32) *Employee {
	p := etype.DeepCopy()
	p.Name = name
	p.Location.Unit = unit

	return p
}
func NewPermanentEmployee(name string, unit int32) *Employee {
	return newEmployee(permEmployee, name, unit)
}

func NewContractEmployee(name string, unit int32) *Employee {
	return newEmployee(contrEmployee, name, unit)
}

func main() {
	prat := NewContractEmployee("prat", 323)
	sam := NewPermanentEmployee("sam", 245)

	fmt.Println(prat, prat.Location)
	fmt.Println(sam, sam.Location)
}
