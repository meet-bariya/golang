package main

import "fmt"

func main() {
	defer fmt.Println("Third Statement")
	defer fmt.Println("Second Statement")
	defer fmt.Println("First Statement")
	fmt.Println("Defer statement follows LIFO")
}
