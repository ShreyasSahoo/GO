package main

import "fmt"

func main() {
	fmt.Println("No inheritance in go and no super or parent class as well")
	user1 := User{"shreyas", "shreyassahoo@gmail.com", 20, true}
	fmt.Printf("%v\n", user1)
	fmt.Printf("%+v\n", user1)
	user1.NewMail("shreyas@gmail.com")
	fmt.Printf("%+v\n", user1)
}

// capital U makes the struct public
type User struct {
	Name      string
	Email     string
	Age       int
	isPremium bool
}

func (user User) GetStatus() bool {
	return user.isPremium
}

// gets the copy of user instance & does not change the

func (user User) NewMail(email string) {
	user.Email = email
	fmt.Printf("%+v\n", user)
}
