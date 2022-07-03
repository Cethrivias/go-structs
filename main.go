package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 88005553535,
		},
	}

	alex.updateName("Jim")

	alex.print()
}

func (p *person) updateName (firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
