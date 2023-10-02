package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our biryani: ")

	//comma ok syntax or error ok syntax

	input, error := reader.ReadString('\n')
	fmt.Println("Thanks for the rating , ", input)
	fmt.Println("error is ", error)
}
