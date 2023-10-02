package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("02-01-2006 Monday 03:04:05PM")) //always use this for format

	createdDate := time.Date(2002, time.October, 26, 10, 10, 10, 0, time.UTC)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))
}
