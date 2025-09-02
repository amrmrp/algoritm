package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, c int
	fmt.Fscan(in, &n, &c)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i])
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i] > x[j]
	})

	for _, v := range x {
		if c == 0 {
			break
		}
		if v <= c {
			continue
		}
		if v >= 2*c {
			c = 0
			break
		}

		c = 2*c - v
	}

	fmt.Println(c)
}
