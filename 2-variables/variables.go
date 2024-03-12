package main

import "fmt"

func main() {
	const test = "test"
	fmt.Println(test)
	var (
		test2 = "test 2"
		test3 = "test 3"
	)
	fmt.Println(test2, test3)

	const (
		test4 = test + " 4"
		test5 = "test 5"
	)

	fmt.Println(test4, test5)
}
