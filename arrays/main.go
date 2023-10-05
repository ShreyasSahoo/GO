package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go")

	var fruitList [4]string //to use array you have to provide its size first

	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[2] = "banana"

	fmt.Println(fruitList)      // [apple orange banana ] the space after banana shows the empty last element of the array
	fmt.Println(len(fruitList)) //predifned length

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)
}
