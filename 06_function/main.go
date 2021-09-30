package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous Function Call")
	}()

	fmt.Println("\nAnonymous function with multiple return value")

	num, str := func() (int, string) {
		return 10, "Golang"
	}()

	fmt.Printf("num = %v \nstr = %v", num, str)
}
