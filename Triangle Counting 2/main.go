package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	count := 0

	for a := 1; a <= n; a++ {
		for b := a; b <= n-a; b++ {
			c := n - a - b
			if b <= c && a+b > c {
				count++
			}
		}
	}

	fmt.Println(count)
}
