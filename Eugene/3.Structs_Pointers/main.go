package main

import (
	"fmt"
)

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
		lastName:  "Golovanov",
		contactInfo: contactInfo{
			email:   "shit@gmail.com",
			zipCode: 92620,
		},
	}

	// fmt.Println(&alex)

	// - Method #1 - Update with pointer
	// alexPointer := &alex // &alex - address in memory of alex
	// alexPointer.updateFirstName("Eugene")

	// - Method #2 - Same thing but with shortcut
	alex.updateFirstName("Eugene")

	alex.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	// *pointerToPerson - take value of given memory in addres &alex
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
