package main

import "fmt"

// can not use the walrus operator outside a method
var jwtToken = 300000

// capital first letter 'L' signifies that the const is public
const LoginToken string = "slkdjfla;fdieoarjkndvkl"

func main() {
	fmt.Println("variables")
	var username string = "shreyas"
	fmt.Printf("Variable is of type: %T \n %s \n", username, username)

	var isLoggedIn bool = true
	fmt.Printf("Variable is of type: %T ", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %d ", smallVal)

	var smallFloat float32 = 255.12431413414123413515
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %d \n", smallVal)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable) // 0
	var anotherString string
	fmt.Println(anotherString) // blank

	// implicit type
	var website = "https://google.com"
	fmt.Println(website)

	//no var style
	//walraus operator
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
}
