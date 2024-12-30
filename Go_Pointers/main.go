package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers study of golang")

	// var ptr *int
	// fmt.Println("Value of ptr is: ", ptr)

	Mynumber := 23

	var ptr = &Mynumber
	fmt.Println("Value of Mynumber is: ", ptr)
	fmt.Println("Value of Mynumber is: ", Mynumber)
	fmt.Println("Value of Mynumber is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of Mynumber is: ", Mynumber)
}
