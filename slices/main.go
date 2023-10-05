package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"apple", "orange", "banana"}
	fmt.Println(fruitList)
	fmt.Printf("%T \n", fruitList)

	fruitList = append(fruitList, "Kiwi", "Watermelon")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	highScores := make([]int, 4, 10) //type, length , capacity
	highScores[0] = 123
	highScores[1] = 13
	highScores[2] = 23
	highScores[3] = 12
	// highScores[4] = 1230 -  this line gives an error as the length is only four
	highScores = append(highScores, 124, 1235, 1245)
	fmt.Println(highScores) // here go allocates a new underlying array
	sort.Ints(highScores)   //sorts in place
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a element from a slice based on its index
	index := 1
	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)
}
