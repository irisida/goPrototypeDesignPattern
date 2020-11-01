package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek in
	fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	res := Employee{}
	_ = d.Decode(&res)
	return &res
}

// head office prototype
var headOffice = Employee{
	"", Address{0, "23 Canary Wharf", "London"},
}

// Luxembourg office prototype
var luxOffice = Employee{
	"", Address{0, "65 Ave JFK, Kirchberg", "Luxembourg"},
}

// Brussels office prototype
var bxlOffice = Employee{
	"", Address{0, "100 Grand Palais Ave, Bruxelles", "Bruxelles"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewHeadOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&headOffice, name, suite)
}

func NewLuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&luxOffice, name, suite)
}

func NewBxlOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&bxlOffice, name, suite)
}

func main() {
	john := NewHeadOfficeEmployee("Johnny 95", 123)
	jean := NewLuxOfficeEmployee("Jean Genie", 3000)
	henk := NewBxlOfficeEmployee("Henk Dieu", 20)

	fmt.Println(john)
	fmt.Println(jean)
	fmt.Println(henk)
}
