package main

import (
	"fmt"
)

func main() {
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{}
	fmt.Println(slice)

	slice = append(slice, 10)
	fmt.Println(slice)
	fmt.Println(cap(slice))
}
