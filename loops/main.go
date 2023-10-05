package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

}
