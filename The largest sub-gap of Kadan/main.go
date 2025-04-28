package main

import (
	"fmt"
)

func main(){
	var n int 
	fmt.Scan(&n)

	arr := make([]int,n)

	for i := range arr {
		fmt.Scan(&arr[i])
	}

	maxSoFar  := arr[0]
	maxEndingHere := arr[0]

	for i := 1; i < n; i++ {
		if maxEndingHere + arr[i] > arr[i] {
			maxEndingHere = maxEndingHere + arr[i]
		} else {
			maxEndingHere = arr[i]
		}
		
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
	}

	fmt.Println(maxSoFar)
}