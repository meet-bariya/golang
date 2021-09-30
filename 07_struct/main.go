package main

import "fmt"

func main() {
	fmt.Println("Alice's Details")
	alice := User{"Alice", "SomeSecurePassword", 18}
	fmt.Println("Username = ", alice.username)
	fmt.Println("Age = ", alice.age)

	fmt.Println("\nBob's Details")
	var bob User
	bob.username = "Bob"
	bob.password = "SomeMoreSecurePassword"
	bob.age = 26

	fmt.Println("Username = ", bob.username)
	fmt.Println("Age = ", bob.age)
}

type User struct {
	username, password string
	age                int
}
