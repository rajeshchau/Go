package main

import "fmt"

func main() {

	fmt.Println("Welcome to mapping")

	language := make(map[string]string)

	language["go"] = "golang"
	language["js"] = "javascript"
	language["py"] = "python"

	fmt.Println(language)

	fmt.Println("the js is ", language["js"])

	delete(language, "js")
	fmt.Println(language)

	for key, value := range language {
		// fmt.Printf("key is %vvalue is %v", key, value)
		fmt.Println("key is ", key, "value is ", value)
	}

}
