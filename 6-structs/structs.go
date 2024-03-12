package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	fmt.Println("dale")
	var u user
	u.name = "Sergio"
	u.age = 1
	fmt.Println(u)
}
