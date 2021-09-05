package main

import "fmt"

func main() {

	//var varName datatype = value
	var num int = 43
	var pi float64 = 3.14
	var name string = "go"
	fmt.Printf("%v - %T\n", num, num)
	fmt.Printf("%v - %T\n", pi, pi)
	fmt.Printf("%v - %T\n", name, name)
}
