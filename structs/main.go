package main

import "fmt"

func main() {
	fmt.Println("No inheritance in go and no super or parent class as well")
	user1 := User{"shreyas", "shreyassahoo@gmail.com", 20, true}
	fmt.Printf("%v\n", user1)
	fmt.Printf("%+v\n", user1)
}

// capital U makes the struct public
type User struct {
	Name      string // captialization makes these properties public
	Email     string
	Age       int
	isPremium bool
}
