package main

import "fmt"

func main() {
	chanel := make(chan string, 2)

	chanel <- "canal"
	chanel <- "canal 2"

	message := <-chanel
	message2 := <-chanel

	fmt.Println(message)
	fmt.Println(message2)

}
