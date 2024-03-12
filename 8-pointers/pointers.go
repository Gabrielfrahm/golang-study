package main

import "fmt"

func main() {
	var number1 int = 10
	var number2 int = 0
	number1 = number2

	number2 = 50
	fmt.Println(number1)

	var number3 *int
	var number4 int = 100

	number3 = &number4
	fmt.Println(*number3, number4)

	number4 = 150

	fmt.Println(*number3, number4)
}
