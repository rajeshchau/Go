package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang !!")

	//days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("%d. %s\n", index, day)
	// }

	roughValue := 1

	for roughValue < 10 {

		if roughValue == 3 {
			goto lco
		}

		if roughValue == 5 {
			roughValue += 1
			continue
		}

		fmt.Println(roughValue)
		roughValue += 1
	}
lco:
	fmt.Println("Loop ended")
}
