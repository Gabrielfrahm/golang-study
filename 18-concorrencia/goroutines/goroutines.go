package main

import (
	"fmt"
	"time"
)

func main() {
	go write("aqui")
	write("aqui 2")
}

func write(text string) {
	for {
		time.Sleep(time.Second)
		fmt.Println(text)
	}
}
