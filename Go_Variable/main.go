package main

import "fmt"

//jwtToken := 3000000 // This will not work outside of function

const LoginToken string = "fadgzerdfsgedxgresdf" // this will work as a constant and can be used in anywhere
//in function variable name should be start with capital for be public.

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint16 = 1216
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var flotydate float32 = 1221.12002150205
	fmt.Println(flotydate)
	fmt.Printf("Variable is of type: %T \n", flotydate)

	var anothervar bool
	fmt.Println(anothervar)
	fmt.Printf("Variable is of type: %T \n", anothervar)

	var website = "learncodeonline.in"
	fmt.Println(website)

	numberofUsers := 300000
	fmt.Println(numberofUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
