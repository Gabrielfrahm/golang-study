package main

import (
	"fmt"
	"module/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("here")
	assistant.Write()

	err := checkmail.ValidateFormat("salamegmail.com")
	fmt.Println(err)
}
