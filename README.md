# goPrototypeDesignPattern

The prototype design pattern in go - The prototype pattern reflects a certain real worldliness about object and `thing` creation. In the real world it is seldom that complex objects are made completely from scratch. It is far more typical for objects to be created from reiterating existing designs, or by composing from existing and extended objects.

In this pattern we take a (partially or fully constructed) design as the prototype. We make a copy and extend it. To do this we require

- `Deep copy` support
- This cloning can be made to a factory.

Here we can see an example of the shallow copying issue and how a deep copy prevents it, bt while this code examples how we can solve the top level issue it is not a scalable solution as in a many person application it would lead to a lot of code.

```go
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
```
