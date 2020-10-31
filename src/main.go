package main

import "fmt"

type Address struct {
	HouseNumber, Street, Area, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func (p *Person) PrintDetails() {
	fmt.Printf("%s lives at %s %s %s in %s, %s \n", p.Name, p.Address.HouseNumber, p.Address.Street, p.Address.Area, p.Address.City, p.Address.Country)
}

func main() {

	// create the base person and address object
	john := Person{"John",
		&Address{"123", "London Road", "Drygate", "Glasgow", "Scotland"},
	}

	// show John
	john.PrintDetails()

	// create jane object from John
	jane := john                     // shallow copy
	jane.Name = "Jane"               // update name
	jane.Address.HouseNumber = "456" // update house number

	jane.PrintDetails()
	john.PrintDetails() // note we see Johns address has changed.

	jimi := john
	jimi.Address = &Address{
		john.Address.HouseNumber,
		john.Address.Street,
		john.Address.Area,
		john.Address.City,
		john.Address.Country,
	}

	jimi.Name = "Jimi"
	jimi.Address.HouseNumber = "789"

	jimi.PrintDetails()
	john.PrintDetails() // note we see Johns address has not changed as we updated Jim's address.

}
