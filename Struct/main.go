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
	//First Style
	alex := person{"Alex", "Anderson", contactInfo{"alexa@gmail.com", 95000}}
	fmt.Println(alex)

	//Second Style
	alex2 := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alexa@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Println(alex2.firstName, alex2.lastName)

	//Third Style
	var alex3 person
	fmt.Println(alex3)
	fmt.Printf("%+v", alex3)

	alex3.firstName = "Alex"
	alex3.lastName = "Anderson"
	alex3.email = "alexa@gmail.com"
	alex3.zipCode = 91111

	alex3.print()
	//alexPointer := &alex3
	alex3.updateName("David")
	alex3.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerPerson *person) updateName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}
