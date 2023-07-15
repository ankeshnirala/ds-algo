package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {

	person := new(Person)

	person.Name = "Ankesh"
	person.Age = 26
	person.Address = "Matihani, Begusarai"

	fmt.Println(*person)

	newPerson := Person{Name: "a", Age: 21, Address: "b"}
	fmt.Println(newPerson)
}
