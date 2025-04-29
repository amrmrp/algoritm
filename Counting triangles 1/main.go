package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for a := 1; a <= n/3; a++ {
		for b := a; b <= (n-a)/2; b++ {
			c := n - a - b
			if c < b {
				continue
			}
			if a+b > c {
				count++
			}
		}
	}

	fmt.Println(count)
}
