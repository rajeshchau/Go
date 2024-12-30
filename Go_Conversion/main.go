package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welocme to my pizza app")
	fmt.Println("Please enter the rating of the pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating the pizza", input)

	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //here we are converting the string to float64
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Thanks for rating the pizza", rating+1)
	}
}
