package main

import (
	"fmt"
)

var cnt = 1

func hanoi(from, to, help rune, n int) {
	if n == 1 {
		fmt.Printf("%d %c %c\n", cnt, from, to)
		cnt++
		return
	}

	hanoi(from, help, to, n-1)
	fmt.Printf("%d %c %c\n", cnt, from, to)
	cnt++
	hanoi(help, to, from, n-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	hanoi('A', 'B', 'C', n)
}
