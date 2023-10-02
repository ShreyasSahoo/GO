package main

import "fmt"

func main() {
	fmt.Println("Pointers Begin!")

	var ptr *int
	fmt.Println(ptr) //nil

	myNum := 26
	var ptr1 *int = &myNum
	fmt.Println(ptr1)  //address of myNum memory
	fmt.Println(*ptr1) //value in the address stored in ptr1

	*ptr1 = *ptr1 + 5
	fmt.Println(myNum)
}
