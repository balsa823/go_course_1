package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Thomas",
		contact: contact{
			email:   "jim@gmail.com",
			zipCode: 1500,
		},
	}
	jim.updateName("Mark")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	(*p).firstname = name
}
