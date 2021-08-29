package main

import "fmt"

func main() {

    var num int = 10
    fmt.Println("num = ", num)
    for i:=1; i<=10; i++{
        fmt.Println(i)
    }

    if num%2==0{
        fmt.Println("Xd")
    }
}