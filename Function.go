package main

import "fmt"

func main() {
   num1 := 89
   num2 := 22
   var ans int

   ans = max(num1, num2)

   fmt.Printf( "Max value is : %v\n", ans )
}

func max(n1, n2 int) int {
   var result int

   if (n1 > n2) {
      result = n1
   } else {
      result = n2
   }
   return result 
}
