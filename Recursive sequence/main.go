package main

import (
	"fmt"
)

var memo = make(map[int]int)

func f(n int) int {

	if val, ok := memo[n]; ok {
		return val
	}

	var result int
	if n == 0 {
		result = 5
	} else {
		tmp := f(n - 1) 
		if n%2 == 0 {
			result = tmp - 21
		} else {
			result = tmp * tmp
		}
	}

	memo[n] = result 
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(f(n))
}
