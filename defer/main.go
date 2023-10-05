package main

import "fmt"

func main() {
	defer fmt.Println("This line will go to the end of this scope")
	defer fmt.Println("one")
	defer fmt.Println("two") //defer works in LIFO manner
	fmt.Println("Defer in GO")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 4 goes in last , then 3 then 2 then 1 then 0
		// 4 comes out 1st and so on
	}
}
