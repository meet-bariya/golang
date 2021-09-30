package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter rating for Pizza")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	rating, _ := strconv.ParseFloat(input, 64)
	fmt.Printf("Thank you for %v star Rating", rating)
}
