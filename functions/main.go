package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	result := add(10, 12)
	fmt.Println(result)

	val, msg := addMultNums(1, 2, 3, 5)
	fmt.Println(msg, val)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
func addMultNums(values ...int) (int, string) { //take in any number of int arguments
	var result int = 0
	for _, value := range values {
		result += value
		fmt.Println(value)
	}
	return result, "Your result is "
}
