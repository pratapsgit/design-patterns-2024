package student

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Street  string
	City    string
	Pincode string
}

type Student struct {
	Name     string
	Location *Address
	Friends  []string
}

func (l *Address) DeepCopy() *Address {
	return &Address{
		l.Street,
		l.City,
		l.Pincode}
}

// func (p *Student) DeepCopy() *Student {
// 	q := *p

// 	q.Location = p.Location.DeepCopy()
// 	copy(q.Friends, p.Friends)

// 	return &q
// }

func (p *Student) DeepCopy() *Student {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := Student{}
	_ = d.Decode(&result)

	return &result
}
