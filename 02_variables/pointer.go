package main

import "fmt"

func main() {
	num := 78
	ptr := &num
	fmt.Printf("%d - %T\n", ptr, ptr)
	fmt.Printf("%d", *ptr)
}
