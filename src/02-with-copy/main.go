package main

import "fmt"

type Address struct {
	HouseNumber, Street, Area, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.HouseNumber,
		a.Street,
		a.Area,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name     string
	Address  *Address
	Contacts []string
}

func (p *Person) DeepCopy() *Person {
	c := *p
	c.Address = p.Address.DeepCopy()
	copy(c.Contacts, p.Contacts)
	return &c
}

func (p *Person) PrintDetails() {
	fmt.Printf("%s lives at %s %s %s in %s, %s \n", p.Name, p.Address.HouseNumber, p.Address.Street, p.Address.Area, p.Address.City, p.Address.Country)
}

func main() {

	// create the base person and address object
	john := Person{"John",
		&Address{"123", "London Road", "Drygate", "Glasgow", "Scotland"},
		[]string{"Renton", "Begbie"},
	}
	// show John
	john.PrintDetails()

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.HouseNumber = "12b"
	jane.Address.Street = "Town Lane"
	jane.Contacts = append(jane.Contacts, "Jimi")

	jane.PrintDetails()
}
