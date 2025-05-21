package main

import(
	"fmt"
	"sort"
)


func main(){

   var n int
   fmt.Scan(&n)
   arr := make([]int,n)
   for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
   }

   sort.Ints(arr)

   for i, val := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
   }
   fmt.Println()
}