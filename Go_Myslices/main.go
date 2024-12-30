package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Go Myslices")

	var fruitList = []string{"Apple", "Banana", "Peach", "Orange"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highscore := make([]int, 4)

	highscore[0] = 123
	highscore[1] = 12
	highscore[2] = 124
	highscore[3] = 125
	//highscore[4] = 123 // Error: index out of range [4] with length 4

	highscore = append(highscore, 11, 22, 33)

	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	var cources = []string{"React", "Angular", "Vue", "Ember", "Svelte"}
	fmt.Println(cources)
	var index int = 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}
