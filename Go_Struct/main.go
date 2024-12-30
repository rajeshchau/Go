package main

import "fmt"

func main() {
	fmt.Println("welcome to Struct in Go")

	hitesh := User{"Rajesh", "Rc8807928@gmail.com", true, 23}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are %+v\n", hitesh)
	fmt.Printf("Name is %v and Email is %v\n", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
