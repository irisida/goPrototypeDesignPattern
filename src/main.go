package main

import "fmt"

type Address struct {
	HouseNumber, Street, Area, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John",
		&Address{"123", "London Road", "Drygate", "Glasgow", "Scotland"},
	}

	// show John
	fmt.Println(john.Name, john.Address)

	// create jane
	jane := john                     // shallow copy
	jane.Name = "Jane"               // update name
	jane.Address.HouseNumber = "456" // update house number

	fmt.Println(jane.Name, jane.Address)
	fmt.Println(john.Name, john.Address)

}
