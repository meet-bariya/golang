package main

import "fmt"

func main(){
	num1:= 10
	num2:= 40
	num3:= 58
	if num1>num2{
		if num1>num3{
			fmt.Printf("%d is max",num1)
		}else{
			fmt.Printf("%d is max",num3)
		}
	}else{
		if num2>num3{
			fmt.Printf("%d is max",num2)
		}else{
			fmt.Printf("%d is max",num3)
		}
	}
}