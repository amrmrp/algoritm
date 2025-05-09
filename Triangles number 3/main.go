package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var answer int64 = 0

	for a := 1; a <= n/3; a++ {
		l := n/2 - 2*a + 1
		if l < 0 {
			l = 0
		}
		r := (n - 3*a) / 2
		if r >= l {
			answer += int64(r - l + 1)
		}
	}

	fmt.Println(answer)
}
