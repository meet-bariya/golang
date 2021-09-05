package main

import "fmt"

func main(){

	arr := [5][]int
        arr = {19,78,55,43,90}

        nums := arr
        nums[2] = 78
	
	fmt.Println(arr)
        fmt.Println(nums)
}
