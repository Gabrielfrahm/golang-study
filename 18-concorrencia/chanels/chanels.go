package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan string)
	go write("canal", chanel)
	// for {
	// 	message, isOpen := <-chanel
	// 	if !isOpen {
	// 		break
	// 	}

	// 	fmt.Println(message)
	// }
	fmt.Printf("%T \n", chanel)
	for message := range chanel {
		fmt.Println(message)
	}
	fmt.Println("five chanels")
}

func write(text string, chanel chan string) {
	for i := 0; i < 5; i++ {
		chanel <- text
		time.Sleep(time.Second)
	}

	close(chanel)
}
