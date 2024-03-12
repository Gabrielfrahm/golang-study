package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int
}

type student struct {
	person
	course  string
	faculty string
}

func main() {
	person1 := person{
		name:     "Gabriel",
		lastName: "Frahm",
		age:      25,
	}

	student1 := student{
		person:  person1,
		course:  "ads",
		faculty: "cruzeiro do sul",
	}

	fmt.Println(student1)
	fmt.Println(student1.lastName)
}
