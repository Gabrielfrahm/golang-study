package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go func() {
		write("aqui")
		waitGroup.Done()
	}()

	go func() {
		write("aqui 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(text)
	}
}
