package main

import (
	"fmt"
	"initial/address"
)

func main() {
	typeAddress := address.TypeOfAddress("Avenida Paulista")
	fmt.Println(typeAddress)
}
