package main

import "fmt"

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// taey := person{
	// 	"Nitipat",
	// 	"L",
	// }

	// taey := person{
	// 	firstName: "Nitipat",
	// 	lastName:  "L",
	// }

	// var taey person
	// taey.firstName = "Nitipat"
	// taey.lastName = "L"

	taey := person{
		firstName: "Nitipat",
		lastName:  "L",
		contactInfo: contactInfo{
			email:   "iamsvz@gmail.com",
			zipCode: 50200,
		},
	}

	// fmt.Println(taey)

	taey.print()
	taey.updateName("Nitipat2")
	taey.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(name string) {
	p.firstName = name
}
