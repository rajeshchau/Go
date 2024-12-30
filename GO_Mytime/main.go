package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	persenttime := time.Now()
	fmt.Println(persenttime)

	fmt.Println(persenttime.Format(time.ANSIC))
	fmt.Println(persenttime.Format(time.RFC1123))
	fmt.Println(persenttime.Format(time.RFC3339))
	fmt.Println(persenttime.Format("2006-01-02 15:04:05 monday"))

	createddate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createddate)
	fmt.Println(createddate.Format("2006-01-02 15:04:05 monday"))

}
