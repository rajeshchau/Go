package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Go !!!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Mango"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg List is: ", vegList)
	fmt.Println("Veg List is: ", len(vegList))

}
